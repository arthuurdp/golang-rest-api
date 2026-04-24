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
	Password string 
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

	user.Name = req.Name
	user.Email = req.Email
	user.Password = req.Password
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