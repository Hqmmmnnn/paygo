package http

import (
	"encoding/json"
	"net/http"

	"github.com/Hqqm/paygo/internal/user/domain/interfaces"
)

// UserService ...
type UserService struct {
	UserUsecases interfaces.UserUsecases
}

// NewHandler ...
func NewUserService(userUC interfaces.UserUsecases) *UserService {
	return &UserService{
		UserUsecases: userUC,
	}
}

type userInput struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
}

// CreateUser ...
func (us *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := &userInput{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err := us.UserUsecases.CreateUser(ctx, user.Email, user.Password, user.FirstName, user.LastName, user.Patronymic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}
