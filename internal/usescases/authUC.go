package usescases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Hqqm/paygo/internal/entities"
	"github.com/Hqqm/paygo/internal/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthClaims struct {
	Account *entities.Account
	jwt.StandardClaims
}

type AuthUsecases struct {
	AccountRepository interfaces.AccountRepository
	signingKey        []byte
	expiresAt         time.Duration
}

func NewAuthUsecases(accRep interfaces.AccountRepository, sk []byte, tokenTTLSeconds time.Duration) interfaces.AuthUsecases {
	return &AuthUsecases{
		AccountRepository: accRep,
		signingKey:        sk,
		expiresAt:         time.Second * tokenTTLSeconds,
	}
}

func (authUC *AuthUsecases) SignUp(ctx context.Context, accountID, email, login, password string) (*entities.Account, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashingPass := string(pwd)

	account := &entities.Account{
		ID:       accountID,
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
		return "", errors.New("incorrect login or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New("incorrect login or password")
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
