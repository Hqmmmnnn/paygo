package usescases

import (
	"context"
	"time"

	"github.com/Hqqm/paygo/internal/user/entities"
	"github.com/Hqqm/paygo/internal/user/interfaces"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type AuthClaims struct {
	User *entities.User `json:"user"`
	jwt.StandardClaims
}

// UserUsecases ...
type userUsecases struct {
	UserRepository interfaces.UserRepository
	signingKey     []byte
	expiresAt      time.Duration
}

func NewUserUsecases(ur interfaces.UserRepository, sk []byte, tokenTTLSeconds time.Duration) interfaces.UserUsecases {
	return &userUsecases{
		UserRepository: ur,
		signingKey:     sk,
		expiresAt:      time.Second * tokenTTLSeconds,
	}
}

// CreateUser ...
func (uc *userUsecases) CreateUser(ctx context.Context, email string, password string, firstName string, lastName string, patronymic string) (*entities.User, error) {
	user := &entities.User{
		ID:         uuid.NewV4(),
		Email:      email,
		Password:   password,
		FirstName:  firstName,
		LastName:   lastName,
		Patronymic: patronymic,
	}

	err := uc.UserRepository.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUsecases) SignIn(ctx context.Context, email string) (string, error) {
	user, err := uc.UserRepository.GetUser(ctx, email)

	if err != nil {
		return "", err
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(uc.expiresAt).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(uc.signingKey)
}
