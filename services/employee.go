package services

import (
	"github.com/jmoiron/sqlx"
	"github.com/vaibhav-ch123/asset-management/database"
	"github.com/vaibhav-ch123/asset-management/errors"
	"github.com/vaibhav-ch123/asset-management/models"
	"github.com/vaibhav-ch123/asset-management/repository"
	"github.com/vaibhav-ch123/asset-management/utils"
)

func RegisterEmployee(employee *models.EmployeeRequest) (string, error) {

	userExist, userExistErr := repository.IsUserExists(employee.Email)

	if userExist {
		return "", errors.ErrEmailExists
	}

	if userExistErr != nil {
		return "", userExistErr
	}

	hashPassword, hashPasswordError := utils.HashPassword(employee.Password)
	if hashPasswordError != nil {
		return "", hashPasswordError
	}
	employee.Password = hashPassword

	var jwtToken string

	txErr := database.Tx(func(tx *sqlx.Tx) error {
		newEmployee, err := repository.RegisterEmployee(employee)
		if err != nil {
			return err
		}

		token, err := utils.CreateJwtToken(newEmployee.ID, newEmployee.EmployeeRole)
		if err != nil {
			return err
		}
		jwtToken = token

		return nil
	})

	if txErr != nil {
		return "", txErr
	}

	return jwtToken, nil
}

func LoginEmployee(email, password string) (string, error) {

	employee, err := repository.GetUserIDAndPasswordByEmail(email)

	if err != nil {
		return "", err
	}

	if err := utils.CheckPassword(password, employee.Password); err != nil {
		return "", errors.ErrPasswordNotMatch
	}

	jwtToken, tokenErr := utils.CreateJwtToken(employee.ID, employee.EmployeeRole)
	if tokenErr != nil {
		return "", tokenErr
	}

	return jwtToken, nil
}

func GetEmployees() ([]models.EmployeeResponse, error) {

	employees, err := repository.GetEmployees()

	if err != nil {
		return nil, err
	}

	return employees, nil
}

func GetEmployee(id string) (*models.EmployeeResponse, error) {

	employee, err := repository.GetEmployeeByID(id)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func UpdateEmployee(employee models.UpdateEmployeeRequest) error {

	if err := repository.UpdateEmployeeByID(employee); err != nil {
		return err
	}

	return nil
}
