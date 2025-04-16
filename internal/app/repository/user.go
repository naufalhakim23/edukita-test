package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"edukita-teaching-grading/internal/app/model"
	"edukita-teaching-grading/internal/pkg"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type (
	IUserRepository interface {
		CreateUser(ctx context.Context, user model.User, tx *sqlx.Tx) (docs model.User, err error)
		GetUserByEmail(ctx context.Context, email string, tx *sqlx.Tx) (docs model.User, err error)
		GetUserByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.User, err error)
		UpdateUserByID(ctx context.Context, user model.User, tx *sqlx.Tx) (docs model.User, err error)
		DeleteUserByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.User, err error)
	}
	UserRepository struct {
		RepositoryOption
	}
)

func InitiateUserRepository(opt RepositoryOption) IUserRepository {
	return &UserRepository{
		RepositoryOption: opt,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User, tx *sqlx.Tx) (docs model.User, err error) {
	query, _, err := goqu.Insert(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_USERS)).
		Rows(user).
		Prepared(true).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowContext(ctx, query).Scan(&docs); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}

	return
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string, tx *sqlx.Tx) (docs model.User, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_USERS)).
		Where(
			goqu.Ex{"email": email},
			goqu.Ex{"is_active": true},
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &docs, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "USER_NOT_FOUND",
				Message:    "user not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("user not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.User, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_USERS)).
		Where(
			goqu.Ex{"id": id},
			goqu.Ex{"is_active": true},
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &docs, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "USER_NOT_FOUND",
				Message:    "user not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("user not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *UserRepository) UpdateUserByID(ctx context.Context, user model.User, tx *sqlx.Tx) (docs model.User, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_USERS)).
		Update().
		Set(user).
		Where(goqu.Ex{"id": user.ID}).
		Prepared(true).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).Scan(&docs); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}

	return
}

func (r *UserRepository) DeleteUserByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.User, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_USERS)).
		Update().
		Set(goqu.Record{"deleted_at": time.Now()}).
		Where(goqu.Ex{"id": id}).
		Prepared(true).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).Scan(&docs); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}

	return
}
