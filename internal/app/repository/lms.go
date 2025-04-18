package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"edukita-teaching-grading/internal/app/model"
	"edukita-teaching-grading/internal/pkg"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type (
	ILearningManagementRepository interface {
		CreateAssignment(ctx context.Context, assignment model.Assignment, tx *sqlx.Tx) (doc model.Assignment, err error)
		GetAssignmentByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Assignment, err error)
		GetAssignmentByTeacherID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Assignment, err error)
		UpdateAssignmentByID(ctx context.Context, assignment model.Assignment, tx *sqlx.Tx) (doc model.Assignment, err error)

		CreateSubmission(ctx context.Context, submission model.Submission, tx *sqlx.Tx) (doc model.Submission, err error)
		GetSubmissionByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Submission, err error)
		GetAllSubmissionsByAssignmentID(ctx context.Context, id string, tx *sqlx.Tx) (docs []model.Submission, err error)
		UpdateSubmissionByID(ctx context.Context, submission model.Submission, tx *sqlx.Tx) (doc model.Submission, err error)
	}
	LearningManagementRepository struct {
		RepositoryOption
	}
)

func InitiateLearningManagementRepository(opt RepositoryOption) ILearningManagementRepository {
	return &LearningManagementRepository{
		RepositoryOption: opt,
	}
}

func (r *LearningManagementRepository) CreateAssignment(ctx context.Context, assignment model.Assignment, tx *sqlx.Tx) (doc model.Assignment, err error) {
	query, _, err := goqu.Insert(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_ASSIGNMENTS)).
		Rows(assignment).
		Prepared(true).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowContext(ctx, query).Scan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}

func (r *LearningManagementRepository) GetAssignmentByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Assignment, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_ASSIGNMENTS)).
		Where(
			goqu.Ex{"id": id},
			goqu.Ex{"is_active": true},
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &doc, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "ASSIGNMENT_NOT_FOUND",
				Message:    "assignment not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("assignment not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}

	return
}

func (r *LearningManagementRepository) GetAssignmentByTeacherID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Assignment, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_ASSIGNMENTS)).
		Where(
			goqu.Ex{"teacher_id": id},
			goqu.Ex{"is_active": true},
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &doc, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "ASSIGNMENT_NOT_FOUND",
				Message:    "assignment not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("assignment not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *LearningManagementRepository) UpdateAssignmentByID(ctx context.Context, assignment model.Assignment, tx *sqlx.Tx) (doc model.Assignment, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_ASSIGNMENTS)).
		Update().
		Set(assignment).
		Where(goqu.Ex{"id": assignment.ID}).
		Prepared(true).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).Scan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}

func (r *LearningManagementRepository) CreateSubmission(ctx context.Context, submission model.Submission, tx *sqlx.Tx) (doc model.Submission, err error) {
	query, _, err := goqu.Insert(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_SUBMISSIONS)).
		Rows(submission).
		Prepared(true).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowContext(ctx, query).Scan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}
func (r *LearningManagementRepository) GetSubmissionByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Submission, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_SUBMISSIONS)).
		Where(
			goqu.Ex{"id": id},
			goqu.Ex{"is_active": true},
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &doc, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "SUBMISSION_NOT_FOUND",
				Message:    "submission not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("submission not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *LearningManagementRepository) GetAllSubmissionsByAssignmentID(ctx context.Context, id string, tx *sqlx.Tx) (docs []model.Submission, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_SUBMISSIONS)).
		Where(
			goqu.Ex{"assignment_id": id},
			goqu.Ex{"is_active": true},
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.Select(&docs, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "SUBMISSION_NOT_FOUND",
				Message:    "submission not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("submission not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *LearningManagementRepository) UpdateSubmissionByID(ctx context.Context, submission model.Submission, tx *sqlx.Tx) (doc model.Submission, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_SUBMISSIONS)).
		Update().
		Set(submission).
		Where(goqu.Ex{"id": submission.ID}).
		Prepared(true).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).Scan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}
