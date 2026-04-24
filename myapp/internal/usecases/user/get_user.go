package user

import (
	"Hello_World/myapp/internal/domain/repositories"
	"context"
	"time"

	"github.com/google/uuid"
)

type GetUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

	return &GetUserResponse{
		ID: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}