package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Hqqm/paygo/internal/entities"
	"github.com/Hqqm/paygo/internal/interfaces"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type accountRepository struct {
	db *sqlx.DB
}

type PostgresAccount struct {
	ID        uuid.UUID      `db:"id"`
	UserID    sql.NullString `db:"user_id"`
	Email     string         `db:"email"`
	Login     string         `db:"login"`
	Password  string         `db:"password"`
	Balance   float64        `db:"balance"`
	CreatedAt time.Time      `db:"created_at"`
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
	postgresAcc := &PostgresAccount{}
	err := accountRepository.db.Get(postgresAcc, "SELECT * FROM accounts WHERE login=$1", login)
	if err != nil {
		return nil, err
	}

	account := accountRepository.convertFromPostgresAccount(postgresAcc)
	return account, nil
}

func (accountRepository *accountRepository) convertFromPostgresAccount(account *PostgresAccount) *entities.Account {
	return &entities.Account{
		ID:        account.ID.String(),
		UserID:    account.UserID.String,
		Email:     account.Email,
		Login:     account.Login,
		Password:  account.Password,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}
}
