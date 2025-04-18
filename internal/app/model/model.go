package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Base contains common fields for all models
type BaseModel struct {
	ID        uuid.UUID  `db:"id"`
	CreatedBy uuid.UUID  `db:"created_by"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedBy *uuid.UUID `db:"updated_by"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedBy *uuid.UUID `db:"deleted_by"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type JWTToken struct {
	jwt.RegisteredClaims
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
