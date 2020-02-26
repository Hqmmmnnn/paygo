package http

import (
	"encoding/json"
	"net/http"

	"github.com/Hqqm/paygo/internal/auth/interfaces"
)

type AuthService struct {
	AuthUsecases interfaces.AuthUsecases
	Middleware   AuthMiddleware
}

func NewAuthService(authUsecases interfaces.AuthUsecases, authMiddleware AuthMiddleware) *AuthService {
	return &AuthService{
		AuthUsecases: authUsecases,
		Middleware:   authMiddleware,
	}
}

type registeringAccount struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (as *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	registerAccount := &registeringAccount{}

	if err := json.NewDecoder(r.Body).Decode(registerAccount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := as.AuthUsecases.SignUp(ctx, registerAccount.Email, registerAccount.Login, registerAccount.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload, err := json.Marshal(account)
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

type signInAccount struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func (as *AuthService) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	signInAccount := &signInAccount{}

	if err := json.NewDecoder(r.Body).Decode(signInAccount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	token, err := as.AuthUsecases.SignIn(ctx, signInAccount.Login, signInAccount.Password)
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
