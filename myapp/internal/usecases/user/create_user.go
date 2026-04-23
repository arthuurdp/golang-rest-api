package user

import (
	"Hello_World/myapp/pkg/apperror"
	"Hello_World/myapp/internal/domain/repositories"
	"Hello_World/myapp/internal/domain/entities"

	"context"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID    string
	Name  string
	Email string
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
		return nil, apperror.NewConflictError("Email already in use")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.NewInternalServerError("Failed to hash password: %v", err)
	}

	user, err := entities.NewUser(req.Name, req.Email, string(hash)) 
	if err != nil {
		return nil, apperror.NewValidationError(err.Error())
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, apperror.NewInternalServerError("User couldn't be created: %w", err)
	}

	return &CreateUserResponse{
		ID: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
	}, nil
}
