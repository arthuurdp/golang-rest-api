package handler

import (
	"net/http"

	userusecase "Hello_World/myapp/internal/usecases/user"
	"Hello_World/myapp/pkg/apperror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	createUser *userusecase.CreateUserUseCase
	updateUser *userusecase.UpdateUserUseCase
	getUser    *userusecase.GetUserUseCase
	getUsers   *userusecase.GetUsersUseCase
}

type createUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type updateUserRequest struct {
	Name     string `json:"name" binding:"omitempty,min=2,max=100"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=8"`
}

func NewUserHandler(create *userusecase.CreateUserUseCase, update *userusecase.UpdateUserUseCase, get *userusecase.GetUserUseCase, getAll *userusecase.GetUsersUseCase) *UserHandler {
	return &UserHandler{createUser: create, updateUser: update, getUser: get, getUsers: getAll}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req createUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, err)
		return
	}

	output, err := h.createUser.Execute(c.Request.Context(), userusecase.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	
	if err != nil {
		apperror.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *UserHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		apperror.HandleError(c, apperror.NewValidationError("invalid id"))
		return
	}

	var req updateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		apperror.HandleError(c, err)
	}

	output, err := h.updateUser.Execute(c.Request.Context(), id, userusecase.UpdateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		apperror.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) FindById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		apperror.HandleError(c, apperror.NewValidationError("invalid id"))
		return
	}

	output, err := h.getUser.Execute(c.Request.Context(), id)
	if err != nil {
		apperror.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *UserHandler) FindAll(c *gin.Context) {
	output, err := h.getUsers.Execute(c.Request.Context())
	if err != nil {
		apperror.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, output)
}
