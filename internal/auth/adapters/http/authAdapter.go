package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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

type SignUpInput struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (as *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	signUpInput := &SignUpInput{}
	if err := json.NewDecoder(r.Body).Decode(signUpInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := as.AuthUsecases.SignUp(ctx, signUpInput.ID, signUpInput.Email, signUpInput.Login, signUpInput.Password)
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

type signInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func (as *AuthService) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	signInInput := &signInInput{}
	if err := json.NewDecoder(r.Body).Decode(signInInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	token, err := as.AuthUsecases.SignIn(ctx, signInInput.Login, signInInput.Password)
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
