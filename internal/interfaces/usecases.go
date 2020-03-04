package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/entities"
)

type AuthUsecases interface {
	SignUp(ctx context.Context, accountID, email, login, password string) (*entities.Account, error)
	SignIn(ctx context.Context, login, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*entities.Account, error)
}
