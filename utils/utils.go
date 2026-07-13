package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/bcrypt"
)

var generator *shortid.Shortid

var generatorSeed int64 = 1000

type clientError struct {
	ID            string `json:"id"`
	MessageToUser string `json:"messageToUser"`
	DeveloperInfo string `json:"developerInfo"`
	Error         string `json:"error"`
	StatusCode    int    `json:"statusCode"`
	IsClientError bool   `json:"isClientError"`
}

func init() {

	n, err := rand.Int(rand.Reader, big.NewInt(generatorSeed))

	if err != nil {
		log.Panicf("failed to initialize utils with random seed: %+v", err)
		return
	}

	g, err := shortid.New(1, shortid.DefaultABC, n.Uint64())
	if err != nil {
		log.Panicf("failed to initialize utils with random seed: %+v", err)
	}

	generator = g
}

func ParseBody(r io.Reader, body any) error {
	if err := json.NewDecoder(r).Decode(body); err != nil {
		return err
	}
	return nil
}

func EncodeJSONBody(w http.ResponseWriter, body interface{}) error {
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		return err
	}
	return nil
}

func ResponseJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	if body != nil {
		if err := EncodeJSONBody(w, body); err != nil {
			logrus.Errorf("failed to response json with error: %+v", err)
		}
	}
}

func newclientError(statusCode int, messageToUser string, err error, additionalInfoDev ...string) *clientError {
	additionalInfo := strings.Join(additionalInfoDev, "/n")

	if additionalInfo == "" {
		additionalInfo = messageToUser
	}

	errorId, err := generator.Generate()

	if err != nil {
		log.Panicf("failed to generate errorId: %+v", err)
	}

	var errString string
	if err != nil {
		errString = err.Error()
	}

	return &clientError{
		errorId,
		messageToUser,
		additionalInfo,
		errString,
		statusCode,
		true,
	}
}

func ResponseError(w http.ResponseWriter, statusCode int, err error, messageToUser string, additionalInfoDev ...string) {
	logrus.Errorf("status: %d, message: %s, err: %+v ", statusCode, messageToUser, err)
	clientError := newclientError(statusCode, messageToUser, err, additionalInfoDev...)
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(clientError); err != nil {
		logrus.Errorf("Failed to send error to caller with error: %+v", err)
	}
}

func CreateJwtToken(employeeID, employeeRole string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"employeeID":   employeeID,
			"employeeRole": employeeRole,
			"ExpiresAt":    time.Now().Add(time.Minute * 5),
			"IssuedAt":     time.Now(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJwtToken(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token.Claims.(jwt.MapClaims), nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}