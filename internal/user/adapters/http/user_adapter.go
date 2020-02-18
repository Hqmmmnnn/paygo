package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Hqqm/paygo/internal/user/interfaces"
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

	res, err := us.UserUsecases.CreateUser(ctx, user.Email, user.Password, user.FirstName, user.LastName, user.Patronymic)
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
	Email string `json:"email"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func (us *UserService) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	signIn := &signInInput{}

	if err := json.NewDecoder(r.Body).Decode(signIn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	token, err := us.UserUsecases.SignIn(ctx, signIn.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	payload, err := json.Marshal(&signInResponse{Token:token})
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}
}
