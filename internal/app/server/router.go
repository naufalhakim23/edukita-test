package server

import (
	"edukita-teaching-grading/internal/app/handler"
	"edukita-teaching-grading/internal/app/server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Router(option handler.HandlerOptions, f *fiber.App) {
	user := handler.UserHandler{HandlerOptions: option}
	lms := handler.LMSHandler{HandlerOptions: option}

	authMiddleware := middlewares.NewAuthMiddleware(option.OptionsApplication)
	v1 := f.Group("/api/v1")

	userGroup := v1.Group("/user")
	userGroup.Post("/register", user.RegisterUser)
	userGroup.Post("/login", user.LoginUser)
	userGroup.Get("/me", authMiddleware.AuthenticateJWT(), user.GetUserByID)
	userGroup.Get("/:id", authMiddleware.AuthenticateJWT(), user.GetUserByID)

	lmsGroup := v1.Group("/lms")
	lmsGroup.Post("/assignments", authMiddleware.AuthenticateJWT(), lms.CreateAssignment)
	lmsGroup.Get("/assignments/:id", authMiddleware.AuthenticateJWT(), lms.GetAssignmentByID)
	lmsGroup.Put("/assignments/:id", authMiddleware.AuthenticateJWT(), lms.UpdateAssignmentByID)

	lmsGroup.Post("/submissions", authMiddleware.AuthenticateJWT(), lms.CreateSubmission)
	lmsGroup.Get("/submissions/:id", authMiddleware.AuthenticateJWT(), lms.GetSubmissionByID)
	lmsGroup.Put("/submissions/:id", authMiddleware.AuthenticateJWT(), lms.UpdateSubmissionByID)

	lmsGroup.Get("/submissions/assignments/:id", authMiddleware.AuthenticateJWT(), lms.GetAllSubmissionsByAssignmentID)
	lmsGroup.Get("/submissions/users/:id", authMiddleware.AuthenticateJWT(), lms.GetAllSubmissionsByUserID)
}
