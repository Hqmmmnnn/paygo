package auth

import (
	"errors"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrUserNotFound = errors.New("user not found")
)