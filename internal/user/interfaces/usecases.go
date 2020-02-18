package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/user/entities"
)

type UserUsecases interface {
	CreateUser(ctx context.Context, email string, password string, firstName string, lastName string, patronymic string) (*entities.User, error)
	SignIn(ctx context.Context, email string) (string, error)
}
