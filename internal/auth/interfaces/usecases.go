package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/auth/entities"
)

type AuthUsecases interface {
	SignUp(ctx context.Context, email string, login string, password string) (*entities.Account, error)
	SignIn(ctx context.Context, login string, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*entities.Account, error)
}
