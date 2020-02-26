package repository

import (
	"context"

	"github.com/Hqqm/paygo/internal/auth/entities"
	"github.com/Hqqm/paygo/internal/auth/interfaces"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type pgUserRepository struct {
	db *sqlx.DB
}

func NewPgUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &pgUserRepository{db: db}
}

func (pg *pgUserRepository) CreateUserID(ctx context.Context, userID uuid.UUID) error {
	query := `	
		INSERT INTO users(id)
		VALUES (:id)
	`

	_, err := pg.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": userID.String(),
	})

	return err
}

func (pg *pgUserRepository) GetUser(ctx context.Context, userID uuid.UUID) (*entities.User, error) {
	user := &entities.User{}

	err := pg.db.Get(user, "SELECT * FROM users WHERE id=$1", userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
