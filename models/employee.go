package models

import "time"

type EmployeeRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
	JoiningDate  string `json:"joiningDate"`
	EmployeeType string `json:"employeeType"`
	EmployeeRole string `json:"employeeRole"`
}
type EmployeeResponse struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Email        string `json:"email" db:"email"`
	Phone        string `json:"phone" db:"phone"`
	JoiningDate  string `json:"joiningDate" db:"joining_date"`
	EmployeeType string `json:"employeeType" db:"employee_type"`
	EmployeeRole string `json:"employeeRole" db:"employee_role"`
}

type UpdateEmployeeRequest struct {
	ID           string  `json:"id"`
	Name         *string `json:"name"`
	Email        *string `json:"email"`
	Password     *string `json:"password"`
	Phone        *string `json:"phone"`
	JoiningDate  *string `json:"joiningDate"`
	EmployeeType *string `json:"employeeType"`
	EmployeeRole *string `json:"employeeRole"`
}

type Employee struct {
	ID           string    `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	Phone        string    `db:"phone"`
	JoiningDate  time.Time `db:"joining_date"`
	EmployeeType string    `db:"employee_type"`
	EmployeeRole string    `db:"employee_role"`
	CreatedAt    time.Time `db:"created_at"`
	UpdateAt     time.Time `db:"updated_at"`
	ArchivedAt   time.Time `db:"archived_at"`
}
