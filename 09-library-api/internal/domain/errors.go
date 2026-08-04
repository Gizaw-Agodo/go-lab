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

	// borrow 
	ErrBorrowNotFound = errors.New("borrow not found")
	ErrBookUnavailable = errors.New("book is not available")
	ErrBorrowLimitReached = errors.New("borrow limit reached")

)