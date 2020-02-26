package repository

import (
	"context"

	"github.com/Hqqm/paygo/internal/auth/entities"
	"github.com/Hqqm/paygo/internal/auth/interfaces"
	"github.com/jmoiron/sqlx"
)

type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) interfaces.AccountRepository {
	return &accountRepository{db: db}
}

func (accountRepository *accountRepository) SaveAccount(ctx context.Context, account *entities.Account) error {
	query := `	
		INSERT INTO accounts(id, user_id, email, login, password)
		VALUES (:id, :user_id, :email, :login, :password)
	`

	_, err := accountRepository.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":       account.ID.String(),
		"user_id":  account.UserID.String(),
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
