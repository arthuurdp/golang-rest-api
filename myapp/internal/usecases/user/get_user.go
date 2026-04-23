package user

import (
	"Hello_World/myapp/internal/domain/repositories"
	"Hello_World/myapp/pkg/apperror"
	"context"
	"time"

	"github.com/google/uuid"
)

type GetUserResponse struct {
	ID uuid.UUID;
	Name string;
	Email string; 
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewGetUserUseCase(userRepo repositories.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{userRepo: userRepo,}
}

func (uc *GetUserUseCase) Execute(ctx context.Context, id uuid.UUID) (*GetUserResponse, error) {
	user, err := uc.userRepo.FindById(ctx, id)

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, apperror.NewResourceNotFoundError("User not found")
	}

	return &GetUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}