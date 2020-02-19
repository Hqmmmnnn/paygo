package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/auth/entities"
)

type AuthUsecases interface {
	SignUp(ctx context.Context, email string, password string, firstName string, lastName string, patronymic string) (*entities.User, error)
	SignIn(ctx context.Context, email, password string) (string, error)
	ParseToken(ctx context.Context, accessToken string) (*entities.User, error)
}
