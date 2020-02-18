package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hqqm/paygo/internal/auth/interfaces"
)

// UserService ...
type AuthService struct {
	AuthUsecases interfaces.AuthUsecases
}

// NewHandler ...
func NewAuthService(authUsecases interfaces.AuthUsecases) *AuthService {
	return &AuthService{
		AuthUsecases: authUsecases,
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
func (as *AuthService) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := &userInput{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := as.AuthUsecases.SignUp(ctx, user.Email, user.Password, user.FirstName, user.LastName, user.Patronymic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	payload, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

}

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func (as *AuthService) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	signIn := &signInInput{}

	if err := json.NewDecoder(r.Body).Decode(signIn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	token, err := as.AuthUsecases.SignIn(ctx, signIn.Email, signIn.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	payload, err := json.Marshal(&signInResponse{Token: token})
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}
