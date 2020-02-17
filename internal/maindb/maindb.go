package maindb

import (
	"github.com/jmoiron/sqlx"
)

// PgUserStorage ...
type pgStorage struct {
	db *sqlx.DB
}

// NewPgUserStorage ...
func NewPgStorage(dsn string) (*pgStorage, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &pgStorage{db: db}, nil
}

func (pg *pgStorage) GetDB() *sqlx.DB {
	return pg.db
}
