package model

import (
	"time"

	"github.com/google/uuid"
)

// Course represents an academic course
type Course struct {
	BaseModel
	Code        string    `db:"code" json:"code"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	StartDate   time.Time `db:"start_date" json:"start_date"`
	EndDate     time.Time `db:"end_date" json:"end_date"`
	IsActive    bool      `db:"is_active" json:"is_active"`
}

// Assignment represents work assigned to students
type Assignment struct {
	BaseModel
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	DueDate     time.Time `db:"due_date" json:"due_date"`
	CourseID    uuid.UUID `db:"course_id" json:"course_id"`
	TeacherID   uuid.UUID `db:"teacher_id" json:"teacher_id"`
	TotalPoints float64   `db:"total_points" json:"total_points"`
	IsPublished bool      `db:"is_published" json:"is_published"`
}

// Submission represents a student's submitted work for an assignment
type Submission struct {
	BaseModel
	AssignmentID uuid.UUID  `db:"assignment_id" json:"assignment_id"`
	StudentID    uuid.UUID  `db:"student_id" json:"student_id"`
	SubmittedAt  time.Time  `db:"submitted_at" json:"submitted_at"`
	Content      string     `db:"content" json:"content"`
	FileURL      string     `db:"file_url" json:"file_url"`
	Grade        *float64   `db:"grade" json:"grade"`
	Feedback     string     `db:"feedback" json:"feedback"`
	GradedAt     *time.Time `db:"graded_at" json:"graded_at"`
	GradedBy     *string    `db:"graded_by" json:"graded_by"`
}
