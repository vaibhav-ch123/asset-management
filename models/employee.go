package models

import "time"

type Employee struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password" db:"password"`
	Phone        string    `json:"phone" db:"phone"`
	JoiningDate  string    `json:"joiningDate" db:"joining_date"`
	EmployeeType string    `json:"employeeType" db:"employee_type"`
	EmployeeRole string    `json:"employeeRole" db:"employee_role"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdateAt     time.Time `json:"updatedAt,omitempty" db:"updated_at"`
	ArchivedAt   time.Time `json:"archivedAt,omitempty" db:"archived_at"`
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