package errors

import "errors"

var (
	ErrEmailExists = errors.New("email already exists!")
	ErrEmailNotExists = errors.New("email not exists!")
	ErrPasswordNotMatch = errors.New("password invalid!")
	ErrEmployeeIDNotMatch = errors.New("user id not found!")
	ErrAssetSerialNumberMatch = errors.New("asset serial number already exist!")
)