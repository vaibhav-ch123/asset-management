package handlers

import (
	
	"net/http"
	"net/mail"
	"strings"
	"time"

	"github.com/vaibhav-ch123/asset-management/errors"
	"github.com/vaibhav-ch123/asset-management/models"
	"github.com/vaibhav-ch123/asset-management/services"
	"github.com/vaibhav-ch123/asset-management/utils"
)

func RegisterEmployee(w http.ResponseWriter, r *http.Request) {

	var body models.Employee

	if err := utils.ParseBody(r.Body, &body); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err, "Failed to parse request body!")
		return
	}

	body.Name = strings.Trim(body.Name, " ")
	body.Email = strings.Trim(body.Email, " ")
	body.Phone = strings.Trim(body.Phone, " ")
	body.Password = strings.Trim(body.Password, " ")

	if len(body.Name) < 3 {
		utils.ResponseError(w, http.StatusBadRequest, nil, "Name Field cannot less then 4!")
		return
	}

	if len(body.Email) > 30 {
		utils.ResponseError(w, http.StatusBadRequest, nil, "mail Field cannot greater then 30!")
		return
	}
    
	if len(body.Password) < 8 {
		utils.ResponseError(w, http.StatusBadRequest, nil, "password Field cannot less then 8!")
		return
	}

	_, mailErr := mail.ParseAddress(body.Email)

	if mailErr != nil {
		utils.ResponseError(w, http.StatusBadRequest, nil, "mail Field format is wrong!")
		return
	}

	if len(body.Phone) != 10 {
		utils.ResponseError(w, http.StatusBadRequest, nil, "phone Field must be equal to 10!")
		return
	}

	for _, ch := range body.Phone {
		if rune('0') > ch || rune('9') < ch {
			utils.ResponseError(w, http.StatusBadRequest, nil, "Phone number is not valid!")
			return
		}
	}

	joiningDate, DateErr := time.Parse("2006-01-01", body.JoiningDate)
	if DateErr != nil {
		utils.ResponseError(w, http.StatusBadRequest, nil, "Joining Date format is not valid!")
		return
	}

	switch body.EmployeeType {
	case "full_time", "intern", "freelancer":
		break
	default:
		utils.ResponseError(w, http.StatusBadRequest, nil, "Employee type is not valid!")
		return
	}

	switch body.EmployeeRole {
	case "admin", "hr", "manager", "developer":
		break
	default:
		utils.ResponseError(w, http.StatusBadRequest, nil, "Employee role is not valid!")
		return
	}

	jwtToken, err := services.RegisterEmployee(&body, joiningDate)

	if err != nil {
		switch err {
		case errors.ErrEmailExists:
			utils.ResponseError(w, http.StatusConflict, err, "Email already exists")
		default:
			utils.ResponseError(w, http.StatusInternalServerError, err, "failed to create employee!")
		}
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, struct {
		JwtToken string `json:"jwtToken"`
	}{
		JwtToken: jwtToken,
	})
}

func LoginEmployee(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := utils.ParseBody(r.Body, &body); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err, "Failed to parse request body!")
		return
	}

	body.Email = strings.Trim(body.Email, " ")
	body.Password = strings.Trim(body.Password, " ")

	jwtToken, err := services.LoginEmployee(body.Email, body.Password)

	if err != nil {
		switch err {
		case errors.ErrEmailNotExists:
			utils.ResponseError(w, http.StatusUnauthorized, err, "Email not exists!")
		case errors.ErrPasswordNotMatch:
			utils.ResponseError(w, http.StatusUnauthorized, err, "Credential invalid!")
		default:
			utils.ResponseError(w, http.StatusInternalServerError, err, "failed to login user!")
		}
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, struct {
		JwtToken string `json:"jwtToken"`
	}{
		JwtToken: jwtToken,
	})
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

    employees, err := services.GetEmployees()

	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err, "failed to fetch employees!")
		return
	}

	utils.ResponseJSON(w, http.StatusOK, employees)
}

