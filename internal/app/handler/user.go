package handler

import (
	"errors"
	"net/http"

	"edukita-teaching-grading/internal/app/payload"
	"edukita-teaching-grading/internal/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	HandlerOptions
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) (err error) {
	var e *pkg.AppError
	req := new(payload.RegisterUserRequest)
	if err = c.BodyParser(req); err != nil {
		return
	}

	v := NewValidator()
	if errs := v.Validate(req); len(errs) > 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request body",
			"errors":  errs,
		})
	}

	res, err := h.Service.User.RegisterUser(c.Context(), *req)
	if err != nil {
		resError := payload.BaseResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Error:   err,
		}
		if errors.As(err, &e) {
			resError.Status = e.StatusCode
			resError.Message = e.Message
			resError.Error = e.Err
		} else {
			resError.Status = http.StatusInternalServerError
		}
		return c.Status(resError.Status).JSON(resError)
	}

	response := payload.BaseResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    res,
	}
	return c.Status(http.StatusOK).JSON(response)
}
