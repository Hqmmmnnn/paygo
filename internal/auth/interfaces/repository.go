package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/auth/entities"
)

type UserRepository interface {
	CreateEmptyUserWithID(ctx context.Context, userID string) error
	GetUser(ctx context.Context, userID string) (*entities.User, error)
}

type AccountRepository interface {
	SaveAccount(ctx context.Context, account *entities.Account) error
	GetAccount(ctx context.Context, email string) (*entities.Account, error)
}
