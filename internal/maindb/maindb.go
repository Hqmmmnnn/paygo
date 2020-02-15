package maindb

import (
	"context"
	"fmt"
	"log"

	"github.com/Hqqm/paygo/internal/domain/entities"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// PgUserStorage ...
type PgUserStorage struct {
	db *sqlx.DB
}

// NewPgUserStorage ...
func NewPgUserStorage() (*PgUserStorage, error) {
	viper.SetConfigFile("/paygo/config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable",
		viper.Get("dbUser"), viper.Get("dbPassword"), viper.Get("dbHost"),
		viper.Get("dbPort"), viper.Get("dbName"))

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

//InitDb initialize postgres database from config.
func InitDb() {
	viper.SetConfigFile("/paygo/config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	dbInfo := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s?sslmode=disable",
		viper.Get("dbUser"), viper.Get("dbPassword"), viper.Get("dbHost"),
		viper.Get("dbPort"), viper.Get("dbName"))

	db, err := sqlx.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
