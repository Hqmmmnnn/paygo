package entities

import (
	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	ID         uuid.UUID `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Patronymic string    `json:"patronymic" db:"patronymic"`
}
