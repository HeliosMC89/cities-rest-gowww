package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(driver, uri string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}
