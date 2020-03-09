package usescases

import (
	"context"

	"github.com/Hqqm/paygo/internal/interfaces"
)

type MoneyOperationsUsecases struct {
	accountRepository interfaces.AccountRepository
}

func NewMoneyOperationsUsecases(accRep interfaces.AccountRepository) *MoneyOperationsUsecases {
	return &MoneyOperationsUsecases{
		accountRepository: accRep,
	}
}

func (moneyOpUC *MoneyOperationsUsecases) ReplenishmentBalance(ctx context.Context, accountID string, amount float64) error {
	err := moneyOpUC.accountRepository.ReplenishmentBalance(ctx, accountID, amount)
	if err != nil {
		return err
	}

	return nil
}
