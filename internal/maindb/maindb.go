package maindb

import (
	"github.com/jmoiron/sqlx"
)

// PgUserStorage ...
type PgStorage struct {
	Db *sqlx.DB
}

// NewPgUserStorage ...
func NewPgStorage(dsn string) (*PgStorage, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PgStorage{Db: db}, nil
}
