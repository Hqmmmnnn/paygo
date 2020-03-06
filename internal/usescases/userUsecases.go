package usescases

import (
	"context"

	"github.com/Hqqm/paygo/internal/entities"
	"github.com/Hqqm/paygo/internal/interfaces"
)

type UserUsecases struct {
	UserRepository    interfaces.UserRepository
	AccountRepository interfaces.AccountRepository
}

func NewUserUsecase(userRep interfaces.UserRepository, accRep interfaces.AccountRepository) interfaces.UserUsecases {
	return &UserUsecases{
		UserRepository:    userRep,
		AccountRepository: accRep,
	}
}

func (userUC *UserUsecases) AddUserInfoToAccount(ctx context.Context, accountID string, user *entities.User) error {
	err := userUC.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	err = userUC.AccountRepository.SetUserID(ctx, accountID, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (userUC *UserUsecases) GetUserById(ctx context.Context, userID string) (*entities.User, error) {
	user, err := userUC.UserRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
