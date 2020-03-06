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

type UserService struct {
	UserUseCases interfaces.UserUsecases
}

func NewAccountService(userUC interfaces.UserUsecases) *UserService {
	return &UserService{
		UserUseCases: userUC,
	}
}

func (userService *UserService) AddUserInfoToAccount(w http.ResponseWriter, r *http.Request) {
	account := r.Context().Value("account").(*entities.Account)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	userInfo := &entities.User{}
	if err := json.NewDecoder(r.Body).Decode(userInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := userService.UserUseCases.AddUserInfoToAccount(ctx, account.ID, userInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type GetUserByIdInput struct {
	ID string `json:"id"`
}

func (userService *UserService) GetUserById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	input := &GetUserByIdInput{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := userService.UserUseCases.GetUserById(ctx, input.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_lib.MarshalJsonAndWrite(user, w)
}
