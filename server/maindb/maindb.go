package maindb

import (
	"github.com/jmoiron/sqlx"
)

type PgStorage struct {
	db *sqlx.DB
}

func NewPgStorage(dsn string) (*PgStorage, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PgStorage{db: db}, nil
}

func (pg *PgStorage) GetDBConnection() *sqlx.DB {
	return pg.db
}
