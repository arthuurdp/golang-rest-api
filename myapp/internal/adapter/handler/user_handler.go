package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"Hello_World/myapp/pkg/apperror"

	userusecase "Hello_World/myapp/internal/usecases/user"
)

type UserHandler struct {
	createUser *userusecase.CreateUserUseCase
	getUser * userusecase.GetUserUseCase
}

func NewUserHandler(create *userusecase.CreateUserUseCase, get *userusecase.GetUserUseCase) *UserHandler {
	return &UserHandler{createUser: create, getUser: get}
}

type createUserRequest struct {
	Name string `json:"name" binding: "required,min=2,max=100"`
	Email string `json:"email" binding: "required,email"`
	Password string `json:"password" binding: "required,min=8"`
}

func (h *UserHandler) Create(c *gin.Context) {
	var req createUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}