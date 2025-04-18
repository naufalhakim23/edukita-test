package service

import (
	"edukita-teaching-grading/internal/app/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTToken(user model.User, secret string, expireTime time.Duration) (string, error) {
	createdAt := user.CreatedAt.Format(time.RFC3339)
	var updatedAt string
	if user.UpdatedAt != nil {
		updatedAt = user.UpdatedAt.Format(time.RFC3339)
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JWTToken{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        user.ID.String(),
			Issuer:    "edukita-teaching-grading",
			ExpiresAt: jwt.NewNumericDate(now.Add(expireTime)),
			IssuedAt:  jwt.NewNumericDate(now),
		},

		Email:     user.Email,
		Name:      user.FirstName,
		Picture:   user.LastName,
		Role:      user.Role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	return token.SignedString([]byte(secret))
}
