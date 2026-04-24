package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field     string    `json:"field"`
	Message   string    `json:"message"`
}

func FormatErrors(err error) []ValidationError {
	var errs []ValidationError

	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, ValidationError{
			Field:     e.Field(),
			Message:   fieldMessage(e),
		})
	}
	return errs
}

func fieldMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s is not valid", e.Field())
	}
}