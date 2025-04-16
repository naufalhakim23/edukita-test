package pkg

import (
	"edukita-teaching-grading/configs"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// wiring all the options
type OptionsApplication struct {
	Config   *configs.Config
	Postgres *sqlx.DB
	Logger   *zap.SugaredLogger
}
