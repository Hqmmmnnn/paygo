package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Account struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Email     string    `json:"email" db:"email"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"password" db:"password"`
	Balance   float64   `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
