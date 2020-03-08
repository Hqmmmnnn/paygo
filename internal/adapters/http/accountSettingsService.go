package http

import (
	"context"
	"encoding/json"
	"errors"
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

func (accSettingsService *AccountSettingsService) GetUserById(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entities.Account)
	if account.UserID == "" {
		http.Error(w, errors.New("user does not exist").Error(), http.StatusNoContent)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	user, err := accSettingsService.ProfileUseCases.GetUserById(ctx, account.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_lib.MarshalJsonAndWrite(user, w)
}
