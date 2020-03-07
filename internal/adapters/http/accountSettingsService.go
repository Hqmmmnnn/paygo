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

type AccountsSettingsService struct {
	ProfileUseCases interfaces.AccountSettingsUsecases
}

func NewAccountsSettingsService(userUC interfaces.AccountSettingsUsecases) *AccountsSettingsService {
	return &AccountsSettingsService{
		ProfileUseCases: userUC,
	}
}

func (accSettingsService *AccountsSettingsService) AddUserInfoToAccount(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entities.Account)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	userInfo := &entities.User{}
	if err := json.NewDecoder(r.Body).Decode(userInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := accSettingsService.ProfileUseCases.AddUserInfoToAccount(ctx, userInfo, account.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type GetUserByIdInput struct {
	ID string `json:"id"`
}

func (accSettingsService *AccountsSettingsService) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	input := &GetUserByIdInput{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := accSettingsService.ProfileUseCases.GetUserById(ctx, input.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_lib.MarshalJsonAndWrite(user, w)
}
