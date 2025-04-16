package server

import (
	"edukita-teaching-grading/internal/app/handler"

	"github.com/gofiber/fiber/v2"
)

func Router(option handler.HandlerOptions, f *fiber.App) {
	user := handler.UserHandler{HandlerOptions: option}
	v1 := f.Group("/api/v1")

	userGroup := v1.Group("/user")
	userGroup.Post("/register", user.RegisterUser)
}
