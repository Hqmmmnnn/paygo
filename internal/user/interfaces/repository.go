package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/user/entities"
)

// UserRepository ..
type UserRepository interface {
	SaveUser(ctx context.Context, user *entities.User) error
}
