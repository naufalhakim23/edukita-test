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
		// Create User General
		CreateUser(ctx context.Context, user model.User, tx *sqlx.Tx) (docs model.User, err error)
		GetUserByEmail(ctx context.Context, email string, tx *sqlx.Tx) (docs model.User, err error)
		GetUserByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.User, err error)
		UpdateUserByID(ctx context.Context, user model.User, tx *sqlx.Tx) (docs model.User, err error)
		DeleteUserByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.User, err error)

		// Create User Teacher
		CreateTeacher(ctx context.Context, teacher model.Teacher, tx *sqlx.Tx) (docs model.Teacher, err error)
		GetTeacherByEmail(ctx context.Context, email string, tx *sqlx.Tx) (docs model.Teacher, err error)
		GetTeacherByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.Teacher, err error)
		UpdateTeacherByID(ctx context.Context, teacher model.Teacher, tx *sqlx.Tx) (docs model.Teacher, err error)
		DeleteTeacherByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.Teacher, err error)

		// Create User Student
		CreateStudent(ctx context.Context, student model.Student, tx *sqlx.Tx) (docs model.Student, err error)
		GetStudentByEmail(ctx context.Context, email string, tx *sqlx.Tx) (docs model.Student, err error)
		GetStudentByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.Student, err error)
		UpdateStudentByID(ctx context.Context, student model.Student, tx *sqlx.Tx) (docs model.Student, err error)
		DeleteStudentByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.Student, err error)
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
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&docs); err != nil {
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
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&docs); err != nil {
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

func (r *UserRepository) CreateTeacher(ctx context.Context, teacher model.Teacher, tx *sqlx.Tx) (docs model.Teacher, err error) {
	query, _, err := goqu.Insert(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_TEACHERS)).
		Rows(teacher).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&docs); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}

	return
}

func (r *UserRepository) GetTeacherByEmail(ctx context.Context, email string, tx *sqlx.Tx) (docs model.Teacher, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_TEACHERS)).
		Where(
			goqu.Ex{"email": email},
			goqu.Ex{"is_active": true},
		).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &docs, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "TEACHER_NOT_FOUND",
				Message:    "teacher not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("teacher not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *UserRepository) GetTeacherByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.Teacher, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_TEACHERS)).
		Where(
			goqu.Ex{"user_id": id},
		).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &docs, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "TEACHER_NOT_FOUND",
				Message:    "teacher not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("teacher not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *UserRepository) UpdateTeacherByID(ctx context.Context, teacher model.Teacher, tx *sqlx.Tx) (docs model.Teacher, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_TEACHERS)).
		Update().
		Set(teacher).
		Where(goqu.Ex{"user_id": teacher.UserID}).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&docs); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}

	return
}

func (r *UserRepository) DeleteTeacherByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.Teacher, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_TEACHERS)).
		Update().
		Set(goqu.Record{"deleted_at": time.Now()}).
		Where(goqu.Ex{"id": id}).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&docs); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}

	return
}

func (r *UserRepository) CreateStudent(ctx context.Context, student model.Student, tx *sqlx.Tx) (docs model.Student, err error) {
	query, _, err := goqu.Insert(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_STUDENTS)).
		Rows(student).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&docs); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}

	return
}

func (r *UserRepository) GetStudentByEmail(ctx context.Context, email string, tx *sqlx.Tx) (docs model.Student, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_STUDENTS)).
		Where(
			goqu.Ex{"email": email},
			goqu.Ex{"is_active": true},
		).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &docs, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "STUDENT_NOT_FOUND",
				Message:    "student not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("student not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *UserRepository) GetStudentByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.Student, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_STUDENTS)).
		Where(
			goqu.Ex{"user_id": id},
		).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &docs, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "STUDENT_NOT_FOUND",
				Message:    "student not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("student not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *UserRepository) UpdateStudentByID(ctx context.Context, student model.Student, tx *sqlx.Tx) (docs model.Student, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_STUDENTS)).
		Update().
		Set(student).
		Where(goqu.Ex{"id": student.UserID}).
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

func (r *UserRepository) DeleteStudentByID(ctx context.Context, id string, tx *sqlx.Tx) (docs model.Student, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_STUDENTS)).
		Update().
		Set(goqu.Record{"deleted_at": time.Now()}).
		Where(goqu.Ex{"id": id}).
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
