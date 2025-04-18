package handler

import (
	"edukita-teaching-grading/internal/app/model"
	"edukita-teaching-grading/internal/app/payload"
	"edukita-teaching-grading/internal/pkg"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type LMSHandler struct {
	HandlerOptions
}

func (h *LMSHandler) CreateAssignment(c *fiber.Ctx) (err error) {
	var (
		claim = c.Locals("mw.auth.claims").(model.JWTToken)
		e     *pkg.AppError
	)
	req := new(payload.CreateAssignmentRequest)
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

	if claim.UUID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.LearningManagement.CreateAssignment(c.Context(), req)
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

func (h *LMSHandler) GetAssignmentByID(c *fiber.Ctx) (err error) {
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

	if claim.UUID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.LearningManagement.GetAssignmentByID(c.Context(), query)
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

func (h *LMSHandler) UpdateAssignmentByID(c *fiber.Ctx) (err error) {
	var (
		claim = c.Locals("mw.auth.claims").(model.JWTToken)
		id    = c.Params("id")
		e     *pkg.AppError
	)
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "id is required",
		},
		)
	}

	req := new(payload.UpdateAssignmentRequest)
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

	if claim.UUID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.LearningManagement.UpdateAssignmentByID(c.Context(), id, req)
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

func (h *LMSHandler) CreateSubmission(c *fiber.Ctx) (err error) {
	var (
		claim = c.Locals("mw.auth.claims").(model.JWTToken)
		e     *pkg.AppError
	)
	req := new(payload.CreateSubmissionRequest)
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

	if claim.UUID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.LearningManagement.CreateSubmission(c.Context(), req.AssignmentID, req)
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

func (h *LMSHandler) GetSubmissionByID(c *fiber.Ctx) (err error) {
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

	if claim.UUID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.LearningManagement.GetSubmissionByID(c.Context(), query)
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

func (h *LMSHandler) UpdateSubmissionByID(c *fiber.Ctx) (err error) {
	var (
		claim = c.Locals("mw.auth.claims").(model.JWTToken)
		id    = c.Params("id")
		e     *pkg.AppError
	)
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "id is required",
		},
		)
	}

	req := new(payload.UpdateSubmissionRequest)
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

	if claim.UUID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.LearningManagement.UpdateSubmissionByID(c.Context(), id, req)
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

func (h *LMSHandler) GetAllSubmissionsByAssignmentID(c *fiber.Ctx) (err error) {
	var (
		claim = c.Locals("mw.auth.claims").(model.JWTToken)
		e     *pkg.AppError
	)
	query := c.Query("assignment_id")
	if query == "" {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "assignment_id is required",
		},
		)
	}

	if claim.UUID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.LearningManagement.GetAllSubmissionsByAssignmentID(c.Context(), query, claim.UUID)
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

func (h *LMSHandler) GetAllSubmissionsByUserID(c *fiber.Ctx) (err error) {
	var (
		claim = c.Locals("mw.auth.claims").(model.JWTToken)
		e     *pkg.AppError
	)
	query := c.Query("user_id")
	if query == "" {
		return c.Status(http.StatusBadRequest).JSON(payload.BaseResponse{
			Status:  http.StatusBadRequest,
			Message: "user_id is required",
		},
		)
	}

	if claim.UUID == "" {
		return c.Status(http.StatusUnauthorized).JSON(payload.BaseResponse{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		},
		)
	}

	res, err := h.Service.LearningManagement.GetAllSubmissionsByUserID(c.Context(), query)
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
