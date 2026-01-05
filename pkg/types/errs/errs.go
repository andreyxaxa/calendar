package errs

import "errors"

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrEventNotFound = errors.New("event not found")
	ErrEmptyResult   = errors.New("empty result")
	ErrAlreadyExists = errors.New("already exists")
)
