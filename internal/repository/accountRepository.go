package repository

import (
	"context"

	"github.com/Hqqm/paygo/internal/entities"
	"github.com/Hqqm/paygo/internal/interfaces"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) interfaces.AccountRepository {
	return &accountRepository{db: db}
}

func (accountRepository *accountRepository) SaveAccount(ctx context.Context, account *entities.Account) error {
	accUUID, err := uuid.FromString(account.ID)
	if err != nil {
		return err
	}

	query := `	
		INSERT INTO accounts(id, email, login, password)
		VALUES (:id, :email, :login, :password)
	`

	_, err = accountRepository.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":       accUUID,
		"email":    account.Email,
		"login":    account.Login,
		"password": account.Password,
	})

	return err
}

func (accountRepository *accountRepository) GetAccount(ctx context.Context, login string) (*entities.Account, error) {
	account := &entities.Account{}
	err := accountRepository.db.Get(account, "SELECT * FROM accounts WHERE login=$1", login)
	if err != nil {
		return nil, err
	}

	return account, nil
}

