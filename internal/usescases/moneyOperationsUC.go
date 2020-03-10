package usescases

import (
	"context"

	"github.com/Hqqm/paygo/internal/interfaces"
	"github.com/jmoiron/sqlx"
)

type MoneyOperationsUsecases struct {
	accountRepository  interfaces.AccountRepository
	transferRepository interfaces.TransferRepository
}

func NewMoneyOperationsUsecases(accRep interfaces.AccountRepository, transferRep interfaces.TransferRepository) *MoneyOperationsUsecases {
	return &MoneyOperationsUsecases{
		accountRepository:  accRep,
		transferRepository: transferRep,
	}
}

func (moneyOpUC *MoneyOperationsUsecases) ReplenishmentBalance(ctx context.Context, accountID string, amount float64) error {
	err := moneyOpUC.accountRepository.ReplenishmentBalance(ctx, accountID, amount)
	if err != nil {
		return err
	}

	return nil
}

func (moneyOpUC *MoneyOperationsUsecases) MoneyTransfer(ctx context.Context, moneyTransferID, senderLogin, recipientLogin, comment string, amount float64) error {
	txErr := moneyOpUC.transferRepository.Transaction(func(tx *sqlx.Tx) (err error) {
		err = moneyOpUC.accountRepository.MoneyTransfer(ctx, tx, senderLogin, recipientLogin, amount)
		if err != nil {
			return err
		}

		err = moneyOpUC.transferRepository.InsertMoneyTransferData(ctx, tx, moneyTransferID, senderLogin, recipientLogin, comment, amount)
		if err != nil {
			return err
		}

		return err
	})

	if txErr != nil {
		return txErr
	}

	return nil
}
