package driver

import (
	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver for database/sql
	"github.com/jmoiron/sqlx"
)

type PostgreSQLOption struct {
	DatabaseName string
	URL          string
}

func NewDatabaseDriver(opt PostgreSQLOption) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", opt.URL)
	if err != nil {
		return nil, err
	}

	// Verify the connection is actually working
	if err = db.Ping(); err != nil {
		db.Close() // Close the connection before returning the error
		return nil, err
	}

	return db, nil
}
