package model

import (
	"net/http"
	"time"

	"edukita-teaching-grading/internal/pkg"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User represents the common attributes for all system users
type User struct {
	BaseModel
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"password_hash" json:"-"`
	FirstName    string     `db:"first_name" json:"first_name"`
	LastName     string     `db:"last_name" json:"last_name"`
	Role         string     `db:"role" json:"role"`
	LastLogin    *time.Time `db:"last_login" json:"last_login"`
	IsActive     bool       `db:"is_active" json:"is_active"`
}

// SetPassword hashes and sets the user's password
func (u *User) SetPassword(password string, cost int) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		err = &pkg.AppError{
			Code:    http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
			Details: []string{
				"please check bcrypt for further details",
			},
			StatusCode: http.StatusInternalServerError,
			Err:        err,
		}
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

// CheckPassword verifies the user's password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}

// Teacher extends the User model with teacher-specific fields
type Teacher struct {
	UserID     uuid.UUID `db:"user_id" json:"user_id"`
	Department string    `db:"department" json:"department"`
	Title      string    `db:"title" json:"title"`
}

// Student extends the User model with student-specific fields
type Student struct {
	UserID         uuid.UUID `db:"user_id" json:"user_id"`
	StudentID      string    `db:"student_id" json:"student_id"`
	EnrollmentYear int       `db:"enrollment_year" json:"enrollment_year"`
	Program        string    `db:"program" json:"program"`
}
