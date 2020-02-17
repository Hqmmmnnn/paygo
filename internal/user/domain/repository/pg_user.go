package repository

import (
	"context"

	"github.com/Hqqm/paygo/internal/user/domain/entities"
	"github.com/Hqqm/paygo/internal/user/domain/interfaces"
	"github.com/jmoiron/sqlx"
)

type PgUserRepository struct {
	db *sqlx.DB
}

// NewPgUserStorage ...
func NewPgUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &PgUserRepository{db: db}
}

// SaveUser ...
func (pg *PgUserRepository) SaveUser(ctx context.Context, user *entities.User) error {
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
