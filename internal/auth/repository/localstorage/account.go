package localstorage

import (
	"context"
	"sync"

	"github.com/Hqqm/paygo/internal/auth"
	"github.com/Hqqm/paygo/internal/auth/entities"
)

type AccountLocalStorage struct {
	accounts map[string]*entities.Account
	mutex    *sync.Mutex
}

func NewAccountLocalStorage() *AccountLocalStorage {
	return &AccountLocalStorage{
		accounts: make(map[string]*entities.Account),
		mutex:    new(sync.Mutex),
	}
}

func (accStorage *AccountLocalStorage) SaveAccount(ctx context.Context, account *entities.Account) error {
	accStorage.mutex.Lock()
	accStorage.accounts[account.ID.String()] = account
	accStorage.mutex.Unlock()

	return nil
}

func (accStorage *AccountLocalStorage) GetAccount(ctx context.Context, email string) (*entities.Account, error) {
	accStorage.mutex.Lock()
	defer accStorage.mutex.Unlock()

	for _, account := range accStorage.accounts {
		if account.Email == email {
			return account, nil
		}
	}

	return nil, auth.ErrAccountNotFound
}
