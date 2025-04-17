package pkg

import "net/http"

type AppError struct {
	Code       string   `json:"code"`
	Message    string   `json:"message"`
	Details    []string `json:"details,omitempty"`
	StatusCode int      `json:"-"` // Not serialized in response
	Err        error    `json:"-"` // Not serialized in response
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Helper constructors
func NewError(code string, msg string, statusCode int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    msg,
		StatusCode: statusCode,
		Err:        err,
	}
}

func NewBadRequestError(msg string, err error) *AppError {
	return NewError(http.StatusText(http.StatusBadRequest), msg, http.StatusBadRequest, err)
}

func NewNotFoundError(msg string, err error) *AppError {
	return NewError(http.StatusText(http.StatusNotFound), msg, http.StatusNotFound, err)
}

func NewDatabaseError(err error) *AppError {
	return NewError(http.StatusText(http.StatusInternalServerError), err.Error(), http.StatusInternalServerError, err)
}
