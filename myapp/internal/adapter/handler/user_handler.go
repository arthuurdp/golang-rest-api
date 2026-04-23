package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"errors"
	userusecase "Hello_World/myapp/internal/usecases/user"
	"Hello_World/myapp/pkg/apperror"
)

type UserHandler struct {
	createUser *userusecase.CreateUserUseCase
	getUser    *userusecase.GetUserUseCase
}

func NewUserHandler(create *userusecase.CreateUserUseCase, get *userusecase.GetUserUseCase) *UserHandler {
	return &UserHandler{createUser: create, getUser: get}
}

type createUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *UserHandler) Create(c *gin.Context) {
	var req createUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.createUser.Execute(c.Request.Context(), userusecase.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *UserHandler) GetById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	output, err := h.getUser.Execute(c.Request.Context(), id)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

func handleError(c *gin.Context, err error) {
	var appErr *apperror.AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.StatusCode(), appErr)
		return
	}

	c.JSON(http.StatusInternalServerError, apperror.NewInternalServerError("erro interno do servidor"))
}
