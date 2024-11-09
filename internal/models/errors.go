package models

import (
	"errors"
)

var (
	ErrNoRecord = errors.New("models: No matching record found")

	ErrInvalidCredentials = errors.New("models: invalid credentials")
	
	ErrDuplicatedEmail = errors.New("models: duplicated email")
)