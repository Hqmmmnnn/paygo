package usescases

import (
	"context"
	"time"

	"github.com/Hqqm/paygo/internal/auth/entities"
	"github.com/Hqqm/paygo/internal/auth/interfaces"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthClaims struct {
	User *entities.User `json:"auth"`
	jwt.StandardClaims
}

type AuthUsecases struct {
	UserRepository interfaces.UserRepository
	signingKey     []byte
	expiresAt      time.Duration
}

func NewAuthUsecases(ur interfaces.UserRepository, sk []byte, tokenTTLSeconds time.Duration) interfaces.AuthUsecases {
	return &AuthUsecases{
		UserRepository: ur,
		signingKey:     sk,
		expiresAt:      time.Second * tokenTTLSeconds,
	}
}

func (ac *AuthUsecases) SignUp(ctx context.Context, email string, password string, firstName string, lastName string, patronymic string) (*entities.User, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	hashingPass := string(pwd)

	user := &entities.User{
		ID:         uuid.NewV4(),
		Email:      email,
		Password:   hashingPass,
		FirstName:  firstName,
		LastName:   lastName,
		Patronymic: patronymic,
	}

	saveErr := ac.UserRepository.SaveUser(ctx, user)
	if saveErr != nil {
		return nil, saveErr
	}

	return user, nil
}

func (ac *AuthUsecases) SignIn(ctx context.Context, email, password string) (string, error) {
	user, err := ac.UserRepository.GetUser(ctx, email)
	if err != nil {
		return "", err
	}

	errWrongCred := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errWrongCred != nil && errWrongCred == bcrypt.ErrMismatchedHashAndPassword {
		return "", errWrongCred
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ac.expiresAt).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(ac.signingKey)
}
