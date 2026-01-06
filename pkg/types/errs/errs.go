package errs

import "errors"

var (
	// ErrUserNotFound -.
	ErrUserNotFound = errors.New("user not found")
	// ErrEventNotFound -.
	ErrEventNotFound = errors.New("event not found")
	// ErrEmptyResult -.
	ErrEmptyResult = errors.New("empty result")
	// ErrAlreadyExists -.
	ErrAlreadyExists = errors.New("already exists")
)
