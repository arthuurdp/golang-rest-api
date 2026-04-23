package apperror

import (
	"net/http"
	"time"
)

type AppError struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) StatusCode() int {
	return e.Status
}

func NewResourceNotFoundError(message string) *AppError {
	return &AppError{
		Code:      "RESOURCE_NOT_FOUND",
		Message:   message,
		Status:    http.StatusNotFound,
		Timestamp: time.Now(),
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Code:      "VALIDATION",
		Message:   message,
		Status:    http.StatusBadRequest,
		Timestamp: time.Now(),
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Code:      "CONFLICT",
		Message:   message,
		Status:    http.StatusConflict,
		Timestamp: time.Now(),
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:      "UNAUTHORIZED",
		Message:   message,
		Status:    http.StatusUnauthorized,
		Timestamp: time.Now(),
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:      "INTERNAL_SERVER_ERROR",
		Message:   message,
		Status:    http.StatusInternalServerError,
		Timestamp: time.Now(),
	}
}