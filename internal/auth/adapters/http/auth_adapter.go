package http

import (
	"encoding/json"
	"net/http"

	"github.com/Hqqm/paygo/internal/auth/interfaces"
)

// UserService ...
type AuthService struct {
	AuthUsecases interfaces.AuthUsecases
	Middleware   AuthMiddleware
}

// NewHandler ...
func NewAuthService(authUsecases interfaces.AuthUsecases, authMiddleware AuthMiddleware) *AuthService {
	return &AuthService{
		AuthUsecases: authUsecases,
		Middleware:   authMiddleware,
	}
}

type registeringUser struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
}

// CreateUser ...
func (as *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := &registeringUser{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := as.AuthUsecases.SignUp(ctx, user.Email, user.Password, user.FirstName, user.LastName, user.Patronymic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
}

type signInUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func (as *AuthService) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	signInUser := &signInUser{}

	if err := json.NewDecoder(r.Body).Decode(signInUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	token, err := as.AuthUsecases.SignIn(ctx, signInUser.Email, signInUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	payload, err := json.Marshal(&signInResponse{Token: token})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}
