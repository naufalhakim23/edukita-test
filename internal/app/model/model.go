package model

import (
	"time"

	"github.com/google/uuid"
)

// Base contains common fields for all models
type BaseModel struct {
	ID        uuid.UUID  `db:"type:uuid;primary_key;default:gen_random_uuid()"`
	CreatedBy string     `db:"created_by"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedBy string     `db:"updated_by"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedBy string     `db:"deleted_by"`
	DeletedAt *time.Time `db:"deleted_at"`
}
