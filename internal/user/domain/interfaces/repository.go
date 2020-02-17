package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/user/domain/entities"
)

// UserRepository ..
type UserRepository interface {
	SaveUser(ctx context.Context, user *entities.User) error
}
