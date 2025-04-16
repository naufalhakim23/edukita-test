package pkg

import "github.com/jmoiron/sqlx"

// wiring all the options
type OptionsApplication struct {
	Postgres *sqlx.DB
}
