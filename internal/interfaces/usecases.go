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

type AccountSettingsUsecases interface {
	AddUserInfoToAccount(ctx context.Context, user *entities.User) error
	GetUserById(ctx context.Context, userID string) (*entities.User, error)
	GetAccountByLogin(ctx context.Context, login string) (*entities.Account, error)
}

type MoneyOperationsUsecases interface {
	ReplenishmentBalance(ctx context.Context, moneyTransferId, recipientLogin string, amount float64) error
	MoneyTransfer(ctx context.Context, moneyTransferID, senderLogin, recipientLogin, comment string, amount float64) error
	GetTransfersHistory(ctx context.Context, login string) (*[]entities.Transfer, error)
	GetTransferById(ctx context.Context, transferId string) (*entities.Transfer, error)
}
