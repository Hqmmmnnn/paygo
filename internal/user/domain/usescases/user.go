package usescases

import (
	"context"

	"github.com/Hqqm/paygo/internal/user/domain/entities"
	"github.com/Hqqm/paygo/internal/user/domain/interfaces"
	uuid "github.com/satori/go.uuid"
)

// UserUsecases ...
type userUsecases struct {
	UserRepository interfaces.UserRepository
}

func NewUserUsecases(ur interfaces.UserRepository) interfaces.UserUsecases {
	return &userUsecases{UserRepository: ur}
}

// CreateUser ...
func (uc *userUsecases) CreateUser(ctx context.Context, email string, password string, firstName string, lastName string, patronymic string) (*entities.User, error) {
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
