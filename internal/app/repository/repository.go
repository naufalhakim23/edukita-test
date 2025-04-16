package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"edukita-teaching-grading/internal/pkg"

	"github.com/jmoiron/sqlx"
)

type RepositoryOption struct {
	pkg.OptionsApplication
}

type Repository struct {
	User IUserRepository
}

func TransactionWrapper(ctx context.Context, db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	if db == nil {
		return &pkg.AppError{
			Code:       "DB_NOT_FOUND",
			Message:    "Database connection not found",
			StatusCode: http.StatusInternalServerError,
			Err:        fmt.Errorf("Database connection not found"),
		}
	}
	opt := sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	}

	tx, err := db.BeginTxx(ctx, &opt)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			err = fmt.Errorf("panic: %v", p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		err = &pkg.AppError{
			Code:       "DB_ERROR",
			Message:    "Database error",
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		return err
	}

	return tx.Commit()
}
