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

func (pg *pgUserRepository) CreateEmptyUserWithID(ctx context.Context, userID string) error {
	query := `
		INSERT INTO users(id)
		VALUES (:id)
	`
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return err
	}

	_, err = pg.db.NamedExecContext(ctx, query, map[string]interface{}{
		"id": userUUID,
	})

	return err
}

func (pg *pgUserRepository) GetUser(ctx context.Context, userID string) (*entities.User, error) {
	user := &entities.User{}
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		return nil, err
	}

	err = pg.db.Get(user, "SELECT * FROM users WHERE id=$1", userUUID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
