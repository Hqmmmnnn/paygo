package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/Hqqm/paygo/internal/domain/usescases"
)

// Handler ...
type Handler struct {
	UserUsecases usescases.UserUsecases
}

// NewHandler ...
func NewHandler(userUC usescases.UserUsecases) *Handler {
	return &Handler{
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
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var user *userInput = &userInput{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err := h.UserUsecases.CreateUser(ctx, user.Email, user.Password, user.FirstName, user.LastName, user.Patronymic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}
