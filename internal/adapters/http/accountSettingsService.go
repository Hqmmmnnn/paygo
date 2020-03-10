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

type AccountSettingsService struct {
	ProfileUseCases interfaces.AccountSettingsUsecases
}

func NewAccountSettingsService(userUC interfaces.AccountSettingsUsecases) *AccountSettingsService {
	return &AccountSettingsService{
		ProfileUseCases: userUC,
	}
}

func (accSettingsService *AccountSettingsService) AddUserInfoToAccount(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entities.Account)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	user := &entities.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = account.ID

	err := accSettingsService.ProfileUseCases.AddUserInfoToAccount(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (accSettingsService *AccountSettingsService) GetUserById(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entities.Account)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	user, err := accSettingsService.ProfileUseCases.GetUserById(ctx, account.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_lib.MarshalJsonAndWrite(user, w)
}

func (accSettingsService *AccountSettingsService) GetAccountByLogin(w http.ResponseWriter, r *http.Request) {
	acc := r.Context().Value("account").(*entities.Account)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	account, err := accSettingsService.ProfileUseCases.GetAccountByLogin(ctx, acc.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_lib.MarshalJsonAndWrite(account, w)
}
