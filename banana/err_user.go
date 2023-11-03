package banana

import "errors"

var (
	UserConflict = errors.New("User existed")
	SignUpFail = errors.New("SignUp Failed")
	GetListUserFail = errors.New("Get list user fail")
	UserNotExist = errors.New("User does not exist")
	UpdateUserFail = errors.New("Update user fail")
)