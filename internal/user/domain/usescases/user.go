package usescases

import (
	"context"

	"github.com/Hqqm/paygo/internal/user/domain/entities"
	"github.com/Hqqm/paygo/internal/user/domain/interfaces"
	uuid "github.com/satori/go.uuid"
)

// UserUsecases ...
type UserUsecases struct {
	UserRepository interfaces.UserRepository
}

func NewUserUsecases(ur interfaces.UserRepository) UserUsecases {
	return UserUsecases{UserRepository: ur}
}

// CreateUser ...
func (uc *UserUsecases) CreateUser(ctx context.Context, email string, password string, firstName string, lastName string, patronymic string) (*entities.User, error) {
	user := &entities.User{
		ID:         uuid.NewV4(),
		Email:      email,
		Password:   password,
		FirstName:  firstName,
		LastName:   lastName,
		Patronymic: patronymic,
	}

	err := uc.UserRepository.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}