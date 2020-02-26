package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/auth/entities"
	uuid "github.com/satori/go.uuid"
)

type UserRepository interface {
	CreateUserID(ctx context.Context, userID uuid.UUID) error
	GetUser(ctx context.Context, userID uuid.UUID) (*entities.User, error)
}

type AccountRepository interface {
	SaveAccount(ctx context.Context, account *entities.Account) error
	GetAccount(ctx context.Context, email string) (*entities.Account, error)
}
