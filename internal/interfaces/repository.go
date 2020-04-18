package interfaces

import (
	"context"

	"github.com/Hqqm/paygo/internal/entities"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	AddUserInfoToAccount(ctx context.Context, user *entities.User) error
	GetUser(ctx context.Context, userID string) (*entities.User, error)
}

type AccountRepository interface {
	SaveAccount(ctx context.Context, account *entities.Account) error
	GetAccount(ctx context.Context, login string) (*entities.Account, error)
	ReplenishmentBalance(ctx context.Context, login string, amount float64) error
	MoneyTransfer(ctx context.Context, tx *sqlx.Tx, senderLogin, recipientLogin string, amount float64) error
}

type TransferRepository interface {
	InsertMoneyTransferData(ctx context.Context, tx *sqlx.Tx, moneyTransferID, senderLogin, recipientLogin, comment string, amount float64) error
	GetTransfers(ctx context.Context, login string) (*[]entities.Transfer, error)
	GetTransferById(ctx context.Context, transferId string) (*entities.Transfer, error)
	GetDbConnection() *sqlx.DB
}
