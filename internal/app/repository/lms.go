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
	ILearningManagementRepository interface {
		CreateCourse(ctx context.Context, course model.Course, tx *sqlx.Tx) (doc model.Course, err error)
		GetCourseByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Course, err error)
		GetCourseByCode(ctx context.Context, code string, tx *sqlx.Tx) (doc model.Course, err error)
		GetAllCourses(ctx context.Context, tx *sqlx.Tx) (docs []model.Course, err error)
		UpdateCourseByID(ctx context.Context, course model.Course, tx *sqlx.Tx) (doc model.Course, err error)
		DeleteCourseByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Course, err error)

		CreateAssignment(ctx context.Context, assignment model.Assignment, tx *sqlx.Tx) (doc model.Assignment, err error)
		GetAssignmentByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Assignment, err error)
		GetAssignmentByTeacherID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Assignment, err error)
		GetAllAssignmentsByCourseID(ctx context.Context, id string, tx *sqlx.Tx) (docs []model.Assignment, err error)
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

func (r *LearningManagementRepository) CreateCourse(ctx context.Context, course model.Course, tx *sqlx.Tx) (doc model.Course, err error) {
	query, _, err := goqu.Insert(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_COURSES)).
		Rows(course).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}

func (r *LearningManagementRepository) GetCourseByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Course, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_COURSES)).
		Where(
			goqu.Ex{"id": id},
			goqu.Ex{"is_active": true},
		).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &doc, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "COURSE_NOT_FOUND",
				Message:    "course not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("course not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}

	return
}

func (r *LearningManagementRepository) GetCourseByCode(ctx context.Context, code string, tx *sqlx.Tx) (doc model.Course, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_COURSES)).
		Where(
			goqu.Ex{"code": code},
			goqu.Ex{"is_active": true},
		).
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.GetContext(ctx, &doc, query); err != nil {
		if err == sql.ErrNoRows {
			err = &pkg.AppError{
				Code:       "COURSE_NOT_FOUND",
				Message:    "course not found",
				StatusCode: http.StatusNotFound,
				Err:        fmt.Errorf("course not found"),
			}
		} else {
			err = pkg.NewDatabaseError(err)
			return
		}
		return
	}
	return
}

func (r *LearningManagementRepository) GetAllCourses(ctx context.Context, tx *sqlx.Tx) (docs []model.Course, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_COURSES)).
		Order(goqu.I("name").Asc()).
		ToSQL()
	if err != nil {
		return
	}

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return
	}

	for rows.Next() {
		row := model.Course{}
		if err = rows.StructScan(&row); err != nil {
			return
		}

		docs = append(docs, row)
	}
	return
}

func (r *LearningManagementRepository) UpdateCourseByID(ctx context.Context, course model.Course, tx *sqlx.Tx) (doc model.Course, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_COURSES)).
		Update().
		Set(course).
		Where(goqu.Ex{"id": course.ID}).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}

func (r *LearningManagementRepository) DeleteCourseByID(ctx context.Context, id string, tx *sqlx.Tx) (doc model.Course, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_COURSES)).
		Update().
		Set(goqu.Record{"deleted_at": time.Now()}).
		Where(goqu.Ex{"id": id}).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}

func (r *LearningManagementRepository) CreateAssignment(ctx context.Context, assignment model.Assignment, tx *sqlx.Tx) (doc model.Assignment, err error) {
	query, _, err := goqu.Insert(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_ASSIGNMENTS)).
		Rows(assignment).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&doc); err != nil {
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
		).
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

func (r *LearningManagementRepository) GetAllAssignmentsByCourseID(ctx context.Context, id string, tx *sqlx.Tx) (docs []model.Assignment, err error) {
	query, _, err := goqu.Select("*").
		From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_ASSIGNMENTS)).
		Where(
			goqu.Ex{"course_id": id},
		).
		ToSQL()
	if err != nil {
		return
	}

	rows, err := tx.QueryxContext(ctx, query)
	if err != nil {
		return
	}

	for rows.Next() {
		row := model.Assignment{}
		if err = rows.StructScan(&row); err != nil {
			return
		}

		docs = append(docs, row)
	}
	return
}

func (r *LearningManagementRepository) UpdateAssignmentByID(ctx context.Context, assignment model.Assignment, tx *sqlx.Tx) (doc model.Assignment, err error) {
	query, _, err := goqu.From(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_ASSIGNMENTS)).
		Update().
		Set(assignment).
		Where(goqu.Ex{"id": assignment.ID}).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}

func (r *LearningManagementRepository) CreateSubmission(ctx context.Context, submission model.Submission, tx *sqlx.Tx) (doc model.Submission, err error) {
	query, _, err := goqu.Insert(fmt.Sprintf("%s.%s", pkg.SCHEMA_NAME, pkg.TABLE_SUBMISSIONS)).
		Rows(submission).
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&doc); err != nil {
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
		).
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
		).
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
		Returning("*").
		ToSQL()
	if err != nil {
		return
	}

	if err = tx.QueryRowxContext(ctx, query).StructScan(&doc); err != nil {
		err = pkg.NewDatabaseError(err)
		return
	}
	return
}
