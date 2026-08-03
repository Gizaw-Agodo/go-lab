package domain

import "errors"

var(
	ErrBookNotFound      = errors.New("book not found")
	ErrBookAlreadyExists = errors.New("book already exists")
	

	// user
	ErrUserNotFound = errors.New("user not found")
	ErrDuplicateEmail = errors.New("email already exists")

	//auth
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidCredentials = errors.New("invalid credentials")

)