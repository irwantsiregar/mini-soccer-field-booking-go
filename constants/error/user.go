package error

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrPasswordIncorrect   = errors.New("incorrect password")
	ErrUsernameExist      = errors.New("username already exist")
	ErrEmailExist      = errors.New("email already exist")
	ErrPasswordDoesNotMatch   = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrEmailExist,
	ErrPasswordDoesNotMatch,
}
