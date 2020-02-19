package repository

import (
	"context"

	"github.com/Hqqm/paygo/internal/auth/entities"
	"github.com/Hqqm/paygo/internal/auth/interfaces"
	"github.com/jmoiron/sqlx"
)

type pgUserRepository struct {
	db *sqlx.DB
}

func NewPgUserRepository(db *sqlx.DB) interfaces.UserRepository {
	return &pgUserRepository{db: db}
}

func (pg *pgUserRepository) SaveUser(ctx context.Context, user *entities.User) error {
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

func (pg *pgUserRepository) GetUser(ctx context.Context, email string) (*entities.User, error) {
	user := &entities.User{}

	err := pg.db.Get(user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
