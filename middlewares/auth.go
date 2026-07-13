package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/vaibhav-ch123/asset-management/utils"
)

type ContextKeys string

const (
	employeeContext ContextKeys = "__employeeContext"
	employeeRoleContext ContextKeys = "__employeeRoleContext"
)

func AuthMiddleWare(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")

		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		jwtToken := strings.TrimPrefix(auth, "Bearer ")

		claims, jwtErr := utils.VerifyJwtToken(jwtToken)
		if jwtErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return 
		}
        
		employeeID, ok := claims["employeeID"].(string)
		if !ok {
			utils.ResponseError(w, http.StatusUnauthorized, nil, "Unauthorized!")
			return
		}

		employeeRole, ok := claims["employeeRole"].(string)
		if !ok {
			utils.ResponseError(w, http.StatusUnauthorized, nil, "Unauthorized!")
			return
		}

		ctx := context.WithValue(r.Context(), employeeContext, employeeID)
		ctx = context.WithValue(ctx, employeeRoleContext, employeeRole)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func EmployeeContext(r *http.Request) string {
    if empID, empErr := r.Context().Value(employeeContext).(string); empErr && empID != "" {
		return empID
	}
	return ""
}

func EmployeeRoleContext(r *http.Request) string {
    if empRole, empErr := r.Context().Value(employeeRoleContext).(string); empErr && empRole != "" {
		return empRole
	}
	return ""
}

func ShouldHaveAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        
		employeeRole := EmployeeRoleContext(r)
		
		if employeeRole != "admin" {
           utils.ResponseError(w, http.StatusUnauthorized, nil, "Admin access required!")
		   return
		}

		next.ServeHTTP(w, r)
	})
}