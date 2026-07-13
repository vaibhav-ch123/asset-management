package repository

import (
	"database/sql"
	"time"

	"github.com/vaibhav-ch123/asset-management/database"
	"github.com/vaibhav-ch123/asset-management/errors"
	"github.com/vaibhav-ch123/asset-management/models"
)

func RegisterEmployee(employee *models.Employee, joiningDate time.Time)(models.Employee, error){
    
	SQL := `INSERT INTO employees (name, email, password, phone, joining_date, employee_type, employee_role)
	       VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, employee_role`
	
	var newEmployee models.Employee	

	err := database.AssetDB.Get(&newEmployee, SQL, employee.Name, employee.Email, employee.Password, employee.Phone, joiningDate, employee.EmployeeType, employee.EmployeeRole)
	
	return newEmployee, err
}

func IsUserExists(email string) (bool, error) {

	SQL := `SELECT id FROM employees WHERE email = TRIM($1) AND archived_at IS NULL`
	var id string
	err := database.AssetDB.Get(&id, SQL, email)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

func GetUserIDAndPasswordByEmail(email string) (*models.Employee, error){

	SQL := `SELECT id, password, employee_role FROM employees
	        WHERE email = $1 AND archived_at IS NULL`

	var employee models.Employee		
	
	err := database.AssetDB.Get(&employee, SQL, email)
	
	if err != nil && err == sql.ErrNoRows {
		return &employee, errors.ErrEmailNotExists
	}

	if err != nil {
		return &employee, err
	}

	return &employee, nil
}

func GetEmployees() ([]models.EmployeeResponse, error) {

	SQL := `SELECT
	            id,
				name,
				email,
				phone,
				joining_date,
				employee_type,
				employee_role 
			FROM employees
			WHERE archived_at IS NULL`

	var employees []models.EmployeeResponse	
	if err := database.AssetDB.Select(&employees, SQL); err != nil{
		return nil, err
	}

	return employees, nil
}