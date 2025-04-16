package handler

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"edukita-teaching-grading/internal/app/repository"
	"edukita-teaching-grading/internal/app/service"
	"edukita-teaching-grading/internal/pkg"

	"github.com/go-playground/validator/v10"
)

type HandlerOptions struct {
	pkg.OptionsApplication
	*service.Service
	*repository.Repository
}

// SimpleValidator provides basic validation functionality
type SimpleValidator struct {
	validate *validator.Validate
}

// New creates a new validator instance
func NewValidator() *SimpleValidator {
	v := validator.New()

	// Register tag name function to use JSON field names in errors
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Register custom validation for non-blank strings
	_ = v.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {
		return strings.TrimSpace(fl.Field().String()) != ""
	})

	return &SimpleValidator{validate: v}
}

// Validate validates a struct and returns any validation errors
func (v *SimpleValidator) Validate(data interface{}) []string {
	var errors []string

	err := v.validate.Struct(data)
	if err == nil {
		return errors
	}

	// Convert validation errors to simple strings
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, formatError(e))
		}
	} else {
		errors = append(errors, err.Error())
	}

	return errors
}

// ValidateJSON validates JSON data against a struct
func (v *SimpleValidator) ValidateJSON(data []byte, structType interface{}) []string {
	// Create a new instance of the provided type
	val := reflect.New(reflect.TypeOf(structType)).Interface()

	// Unmarshal JSON into the struct
	if err := json.Unmarshal(data, val); err != nil {
		return []string{fmt.Sprintf("Invalid JSON: %s", err.Error())}
	}

	// Validate the struct
	return v.Validate(val)
}

// formatError converts a validation error to a user-friendly message
func formatError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email", e.Field())
	case "min":
		if e.Type().Kind() == reflect.String {
			return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
		}
		return fmt.Sprintf("%s must be at least %s", e.Field(), e.Param())
	case "max":
		if e.Type().Kind() == reflect.String {
			return fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
		}
		return fmt.Sprintf("%s must be at most %s", e.Field(), e.Param())
	case "notblank":
		return fmt.Sprintf("%s cannot be blank", e.Field())
	default:
		return fmt.Sprintf("%s failed on %s validation", e.Field(), e.Tag())
	}
}
