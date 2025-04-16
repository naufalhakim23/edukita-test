package driver

import "github.com/jmoiron/sqlx"

type PostgreSQLOption struct {
	DatabaseName string
	URL          string
}

func NewDatabaseDriver(opt PostgreSQLOption) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", opt.URL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
