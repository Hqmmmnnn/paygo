package entities

import (
	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	ID         uuid.UUID
	Email      string
	Password   string
	FirstName  string
	LastName   string
	Patronymic string
}
