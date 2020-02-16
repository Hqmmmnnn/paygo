package maindb

import (
	"context"

	"github.com/Hqqm/paygo/internal/domain/entities"
	"github.com/jmoiron/sqlx"
)

// PgUserStorage ...
type PgUserStorage struct {
	db *sqlx.DB
}

// NewPgUserStorage ...
func NewPgUserStorage(dsn string) (*PgUserStorage, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PgUserStorage{db: db}, nil
}

// SaveUser ...
func (pg *PgUserStorage) SaveUser(ctx context.Context, user *entities.User) error {
	query := `
		INSERT INTO users(id, email, password, first_name, last_name, patronymic)
		VALUES (:id, :email, :password, :first_name, :last_name, :patronymic)
	`

	_, err := pg.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id":         user.ID.String(),
		"email":      user.Email,
		"password":   user.Password,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"patronymic": user.Patronymic,
	})

	return err
}
