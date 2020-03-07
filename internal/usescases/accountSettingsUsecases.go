package usescases

import (
	"context"

	"github.com/Hqqm/paygo/internal/entities"
	"github.com/Hqqm/paygo/internal/interfaces"
)

type AccountSettingsUsecases struct {
	UserRepository    interfaces.UserRepository
	AccountRepository interfaces.AccountRepository
}

func NewAccountSettingsUsecases(userRep interfaces.UserRepository, accRep interfaces.AccountRepository) interfaces.AccountSettingsUsecases {
	return &AccountSettingsUsecases{
		UserRepository:    userRep,
		AccountRepository: accRep,
	}
}

func (accSettingsUC *AccountSettingsUsecases) AddUserInfoToAccount(ctx context.Context, user *entities.User, accountID string) error {
	err := accSettingsUC.UserRepository.AddUserInfoToAccount(ctx, user, accountID)
	if err != nil {
		return err
	}

	return nil
}

func (accSettingsUC *AccountSettingsUsecases) GetUserById(ctx context.Context, userID string) (*entities.User, error) {
	user, err := accSettingsUC.UserRepository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
