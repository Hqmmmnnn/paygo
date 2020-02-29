package localstorage

import (
	"context"
	"testing"
	"time"

	"github.com/Hqqm/paygo/internal/auth"
	"github.com/Hqqm/paygo/internal/auth/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	accStorage := NewAccountLocalStorage()

	accId := uuid.NewV4()
	userId := uuid.NewV4()
	now := time.Now()

	newAccount := &entities.Account{
		ID:        accId,
		UserID:    userId,
		Email:     "hqqmJuly@gmail.com",
		Login:     "hqqm",
		Password:  "hqqm1234",
		Balance:   0,
		CreatedAt: now,
	}

	err := accStorage.SaveAccount(context.Background(), newAccount)
	assert.NoError(t, err)

	account, err := accStorage.GetAccount(context.Background(), "hqqmJuly@gmail.com")
	assert.NoError(t, err)
	assert.Equal(t, newAccount, account)

	account, err = accStorage.GetAccount(context.Background(), "blablabla")
	assert.Error(t, err)
	assert.Equal(t, err, auth.ErrUserNotFound)
}
