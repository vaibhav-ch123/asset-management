package errors

import "errors"

var (
	ErrEmailExists = errors.New("email already exists!")
	ErrEmailNotExists = errors.New("email not exists!")
	ErrPasswordNotMatch = errors.New("password invalid!")
)