package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Hqqm/paygo/internal/entities"
	"github.com/Hqqm/paygo/internal/interfaces"
)

type MoneyOperationsService struct {
	MoneyOperationsUC interfaces.MoneyOperationsUsecases
}

func NewMoneyOperationsService(moneyOpUC interfaces.MoneyOperationsUsecases) *MoneyOperationsService {
	return &MoneyOperationsService{
		MoneyOperationsUC: moneyOpUC,
	}
}

type ReplenishmentBalanceInput struct {
	Amount float64 `json:"amount"`
}

func (moneyOpService *MoneyOperationsService) ReplenishmentBalance(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entities.Account)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	balanceInput := &ReplenishmentBalanceInput{}
	if err := json.NewDecoder(r.Body).Decode(balanceInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := moneyOpService.MoneyOperationsUC.ReplenishmentBalance(ctx, account.ID, balanceInput.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
