package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, userID string) (*entities.User, error)
}

type AccountRepository interface {
	SaveAccount(ctx context.Context, account *entities.Account) error
	GetAccount(ctx context.Context, login string) (*entities.Account, error)
}
