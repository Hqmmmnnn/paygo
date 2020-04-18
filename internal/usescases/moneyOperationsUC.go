package usescases

import (
	"context"

	"github.com/Hqqm/paygo/internal/_lib"
	"github.com/Hqqm/paygo/internal/entities"
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

func (moneyOpUC *MoneyOperationsUsecases) ReplenishmentBalance(ctx context.Context, moneyTransferId, login string, amount float64) error {
	dbConnection := moneyOpUC.transferRepository.GetDbConnection()

	txErr := _lib.WithTransaction(dbConnection, func(tx *sqlx.Tx) (err error) {
		err = moneyOpUC.accountRepository.ReplenishmentBalance(ctx, login, amount)
		if err != nil {
			return err
		}

		err = moneyOpUC.transferRepository.InsertMoneyTransferData(ctx, tx, moneyTransferId, "Paygo", login, "replenished balance via Paygo", amount)
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

func (moneyOpUC *MoneyOperationsUsecases) MoneyTransfer(ctx context.Context, moneyTransferID, senderLogin, recipientLogin, comment string, amount float64) error {
	dbConnection := moneyOpUC.transferRepository.GetDbConnection()

	txErr := _lib.WithTransaction(dbConnection, func(tx *sqlx.Tx) (err error) {
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

func (moneyOpUC *MoneyOperationsUsecases) GetTransfersHistory(ctx context.Context, login string) (*[]entities.Transfer, error) {
	transfers, err := moneyOpUC.transferRepository.GetTransfers(ctx, login)

	if err != nil {
		return nil, err
	}

	return transfers, nil
}

func (moneyOpUC *MoneyOperationsUsecases) GetTransferById(ctx context.Context, transferId string) (*entities.Transfer, error) {
	transfer, err := moneyOpUC.transferRepository.GetTransferById(ctx, transferId)

	if err != nil {
		return nil, err
	}

	return transfer, err
}
