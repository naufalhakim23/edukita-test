package middlewares

import (
	"edukita-teaching-grading/internal/app/model"
	"edukita-teaching-grading/internal/app/payload"
	"edukita-teaching-grading/internal/pkg"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	Secret string
	pkg.OptionsApplication
}

func NewAuthMiddleware(optionsApp pkg.OptionsApplication) AuthMiddleware {
	return AuthMiddleware{
		Secret:             optionsApp.Config.Application.Secret,
		OptionsApplication: optionsApp,
	}
}

func (m *AuthMiddleware) AuthenticateJWT() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		tokenString, err := m.extractTokenFromHeader(c)
		if err != nil {
			m.Logger.Warnf(fmt.Sprintf("failed to extract token from header: %s", err.Error()), zap.Error(err))
			return c.Status(fiber.StatusUnauthorized).JSON(payload.BaseResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "invalid token",
			})
		}

		tokenClaims, err := m.extractClaims(tokenString)
		if err != nil {
			m.Logger.Warnf(fmt.Sprintf("failed to extract claims from token: %s", err.Error()), zap.Error(err))
			return c.Status(fiber.StatusUnauthorized).JSON(payload.BaseResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "invalid token",
			})
		}

		if !tokenClaims.Valid {
			m.Logger.Warnf("invalid token", zap.Error(err))
			return c.Status(fiber.StatusUnauthorized).JSON(payload.BaseResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "invalid token",
			})
		}

		claims := tokenClaims.Claims.(jwt.MapClaims)
		myClaims, err := claimToModelJWTToken(claims)
		if err != nil {
			m.Logger.Warnf(fmt.Sprintf("failed to convert claims to model: %s", err.Error()), zap.Error(err))
			return c.Status(fiber.StatusUnauthorized).JSON(payload.BaseResponse{
				Status:  fiber.StatusUnauthorized,
				Message: "invalid token",
			})
		}

		c.Locals("mw.auth.claims", myClaims)
		return c.Next()
	}
}

func (m *AuthMiddleware) extractTokenFromHeader(c *fiber.Ctx) (tokenString string, err error) {
	authHeader := c.Get("Authorization")
	if len(authHeader) < 1 || authHeader == "" {
		authHeader, err = m.GetTokenFromCookie(c)
		if err != nil {
			return "", err
		}
	} else {
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) < 2 {
			return "", fmt.Errorf("invalid token")
		}
		authHeader = arrToken[1]
	}

	return authHeader, nil
}

func (m *AuthMiddleware) GetTokenFromCookie(c *fiber.Ctx) (tokenString string, err error) {
	env := m.Config.Application.Env

	tokenKey := "edukita_lms"
	if env == "uat" {
		tokenKey = "uat_edukita_lms"
	}

	if env == "staging" || env == "local" {
		tokenString = c.Get("Authorization")
		tokenKey = "stg_edukita_lms"
	}
	if tokenString == "" {
		token := c.Cookies(tokenKey)
		if token == "" {
			return "", fmt.Errorf("cookie %s not found", tokenKey)
		}
		// adding "Bearer " so I dont have to change or create new JWT check function
		tokenString = fmt.Sprintf("Bearer %s", token)
	}

	return tokenString, err
}

func (m *AuthMiddleware) extractClaims(tokenString string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	cleanedClaims, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			return nil, fmt.Errorf("unexpected alg value: %v", token.Header["alg"])
		}
		return []byte(m.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return cleanedClaims, nil
}

func claimToModelJWTToken(claims jwt.MapClaims) (model.JWTToken, error) {
	myClaim := model.JWTToken{}
	errMarshall := mapstructure.Decode(claims, &myClaim)
	if errMarshall != nil {
		return myClaim, errMarshall
	}

	return myClaim, nil
}
