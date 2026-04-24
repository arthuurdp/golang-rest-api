package user

import (
	"Hello_World/myapp/internal/domain/entities"
	"Hello_World/myapp/internal/domain/repositories"
	"Hello_World/myapp/pkg/apperror"
	"fmt"

	"context"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	ID       string	
	Name     string	
	Email    string	
	Password string	
}

type CreateUserResponse struct {
	ID    string	`json:"id"`
	Name  string	`json:"name"`
	Email string	`json:"email"`
}

type CreateUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewCreateUserUseCase(userRepo repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	existingUser, _ := uc.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, apperror.NewConflictError("email already in use")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperror.NewInternalServerError(fmt.Sprintf("failed to hash password: %v", err))
	}

	user, err := entities.NewUser(req.Name, req.Email, string(hash)) 
	if err != nil {
		return nil, apperror.NewValidationError(err.Error())
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, apperror.NewInternalServerError(fmt.Sprintf("failed to create user: %v", err))
	}

	return &CreateUserResponse{
		ID: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
	}, nil
}
