package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Hqqm/paygo/internal/_lib"
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
	ID     string  `json:"id"`
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

	err := moneyOpService.MoneyOperationsUC.ReplenishmentBalance(ctx, balanceInput.ID, account.Login, balanceInput.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type MoneyTransferData struct {
	ID             string  `json:"id" `
	RecipientLogin string  `json:"recipient_login" `
	Comment        string  `json:"comment"`
	Amount         float64 `json:"amount" `
}

func (moneyOpService *MoneyOperationsService) MoneyTransfer(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entities.Account)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	moneyTransferData := &MoneyTransferData{}
	if err := json.NewDecoder(r.Body).Decode(moneyTransferData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := moneyOpService.MoneyOperationsUC.MoneyTransfer(
		ctx,
		moneyTransferData.ID,
		account.Login,
		moneyTransferData.RecipientLogin,
		moneyTransferData.Comment,
		moneyTransferData.Amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func (moneyOpService *MoneyOperationsService) GetTransfersHistory(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entities.Account)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	transfers, err := moneyOpService.MoneyOperationsUC.GetTransfersHistory(ctx, account.Login)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_lib.MarshalJsonAndWrite(transfers, w)
}
