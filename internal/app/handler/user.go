package handler

import (
	"errors"
	"net/http"
	"time"

	"edukita-teaching-grading/internal/app/model"
	"edukita-teaching-grading/internal/app/payload"
	"edukita-teaching-grading/internal/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	HandlerOptions
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) (err error) {
	var (
		e *pkg.AppError
	)
	req := new(payload.RegisterUserRequest)
	if err = c.BodyParser(req); err != nil {
		return
	}

	v := NewValidator()
	if errs := v.Validate(req); len(errs) > 0 {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Error:   errs,
		},
		)
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

func (h *UserHandler) LoginUser(c *fiber.Ctx) (err error) {
	var (
		e *pkg.AppError
	)
	req := new(payload.LoginUserRequest)
	if err = c.BodyParser(req); err != nil {
		return
	}

	v := NewValidator()
	if errs := v.Validate(req); len(errs) > 0 {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid request body",
			Error:   errs,
		},
		)
	}

	res, err := h.Service.User.LoginUser(c.Context(), req)
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

	setCookie := fiber.Cookie{
		Name:     h.Config.Cookies.AccessToken,
		Value:    res.Token,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(h.Config.Cookies.SSOExpired),
		Secure:   false,
		HTTPOnly: false,
		SameSite: "Lax",
	}

	c.Cookie(&setCookie)
	return c.Status(http.StatusOK).JSON(response)
}

func (h *UserHandler) GetUserByEmail(c *fiber.Ctx) (err error) {
	var (
		claim = c.Locals("mw.auth.claims").(model.JWTToken)
		e     *pkg.AppError
	)
	email := c.Query("email")
	if email == "" {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "email is required",
		},
		)
	}

	if claim.ID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.User.GetUserByEmail(c.Context(), email)
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

func (h *UserHandler) GetUserByID(c *fiber.Ctx) (err error) {
	var (
		claim = c.Locals("mw.auth.claims").(model.JWTToken)
		e     *pkg.AppError
	)
	query := c.Query("id")
	if query == "" {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "id is required",
		},
		)
	}
	if claim.ID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.User.GetUserByID(c.Context(), query)
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
