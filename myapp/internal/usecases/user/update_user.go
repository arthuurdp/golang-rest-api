package user

import (
	"Hello_World/myapp/internal/domain/repositories"
	"Hello_World/myapp/pkg/apperror"
	"context"
	"time"

	"github.com/google/uuid"
)

type UpdateUserRequest struct {
	Name     string 
	Email    string 
}

type UpdateUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUpdateUserUseCase(userRepo repositories.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{userRepo: userRepo}
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, id uuid.UUID, req UpdateUserRequest) (*UpdateUserResponse, error) {
	user, err := uc.userRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name == "" && req.Email == "" {
		return nil, apperror.NewValidationError("at least one field must be provided")
	}

	if req.Email != "" {
		existingUser, err := uc.userRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			return nil, err
		}
		if existingUser != nil && existingUser.Email == req.Email {
			return nil, apperror.NewConflictError("email already in use")
		} 
		
		user.Email = req.Email
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	user.UpdatedAt = time.Now()

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, apperror.NewInternalServerError("error updating user")
	}
	
	return &UpdateUserResponse{
		ID: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
		UpdatedAt: user.UpdatedAt,
	}, nil
}