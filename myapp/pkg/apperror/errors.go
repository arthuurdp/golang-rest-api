package apperror

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	appvalidator "Hello_World/myapp/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AppError struct {
	Timestamp time.Time `json:"timestamp"`
	Code      int       `json:"code"`
	Status    string    `json:"status"`
	Message   string    `json:"message"`
}

type validationResponse struct {
	Timestamp time.Time                      `json:"timestamp"`
	Code      int                            `json:"code"`
	Status    string                         `json:"status"`
	Message   string                         `json:"message"`
	Errors    []appvalidator.ValidationError `json:"errors"`
}

func (e *AppError) Error() string  { return e.Message }
func (e *AppError) StatusCode() int { return e.Code }

func HandleError(c *gin.Context, err error) {
	if errors.Is(err, io.EOF) {
		c.JSON(http.StatusBadRequest, AppError{
			Timestamp: time.Now(),
			Code:      http.StatusBadRequest,
			Status:    "BAD_REQUEST",
			Message:   "request body cannot be empty",
		})
		return
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		c.JSON(http.StatusBadRequest, AppError{
			Timestamp: time.Now(),
			Code:      http.StatusBadRequest,
			Status:    "BAD_REQUEST",
			Message:   "malformed JSON",
		})
		return
	}

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		c.JSON(http.StatusBadRequest, validationResponse{
			Timestamp: time.Now(),
			Code:      http.StatusBadRequest,
			Status:    "VALIDATION",
			Message:   "validation failed",
			Errors:    appvalidator.FormatErrors(err),
		})
		return
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.StatusCode(), appErr)
		return
	}

	c.JSON(http.StatusInternalServerError, NewInternalServerError("Internal server error"))
}

func NewResourceNotFoundError(message string) *AppError {
	return &AppError{
		Timestamp: time.Now(), 
		Code: http.StatusNotFound, 
		Status: "NOT_FOUND",
	    Message: message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Timestamp: time.Now(),
		Code:      http.StatusBadRequest,
		Status:    "BAD_REQUEST",
		Message:   message,
	}
}

func NewConflictError(message string) *AppError {
	return &AppError{
		Timestamp: time.Now(),
		Code:      http.StatusConflict,
		Status:    "CONFLICT",
		Message:   message,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Timestamp: time.Now(),
		Code:      http.StatusUnauthorized,
		Status:    "UNAUTHORIZED",
		Message:   message,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Timestamp: time.Now(),
		Code:      http.StatusInternalServerError,
		Status:    "INTERNAL_SERVER_ERROR",
		Message:   message,
	}
}