package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

const postgres = "postgres"

// New database connection
func New(sqlConnection string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(postgres, sqlConnection)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}
	return db, nil
}