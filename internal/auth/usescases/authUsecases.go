package usescases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Hqqm/paygo/internal/auth/entities"
	"github.com/Hqqm/paygo/internal/auth/interfaces"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthClaims struct {
	Account *entities.Account `json:"auth"`
	jwt.StandardClaims
}

type AuthUsecases struct {
	AccountRepository interfaces.AccountRepository
	UserRepository    interfaces.UserRepository
	signingKey        []byte
	expiresAt         time.Duration
}

func NewAuthUsecases(accRep interfaces.AccountRepository, userRep interfaces.UserRepository, sk []byte, tokenTTLSeconds time.Duration) interfaces.AuthUsecases {
	return &AuthUsecases{
		AccountRepository: accRep,
		UserRepository:    userRep,
		signingKey:        sk,
		expiresAt:         time.Second * tokenTTLSeconds,
	}
}

func (authUC *AuthUsecases) SignUp(ctx context.Context, email string, login string, password string) (*entities.Account, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashingPass := string(pwd)

	userID := uuid.NewV4()
	err = authUC.UserRepository.CreateUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	account := &entities.Account{
		ID:       uuid.NewV4(),
		UserID:   userID,
		Email:    email,
		Login:    login,
		Password: hashingPass,
	}

	err = authUC.AccountRepository.SaveAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (authUC *AuthUsecases) SignIn(ctx context.Context, login, password string) (string, error) {
	account, err := authUC.AccountRepository.GetAccount(ctx, login)
	if err != nil {
		return "", err
	}

	errWrongCred := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if errWrongCred != nil && errWrongCred == bcrypt.ErrMismatchedHashAndPassword {
		return "", errWrongCred
	}

	claims := AuthClaims{
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(authUC.expiresAt).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(authUC.signingKey)
}

func (authUC *AuthUsecases) ParseToken(ctx context.Context, accessToken string) (*entities.Account, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["error"])
		}
		return authUC.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.Account, nil
	}

	return nil, errors.New("invalid access token")
}
