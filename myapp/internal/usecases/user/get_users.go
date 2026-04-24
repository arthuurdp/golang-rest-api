package user

import (
	"Hello_World/myapp/internal/domain/repositories"
	"context"
	"time"
)

type GetUsersResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetUsersUseCase struct {
	userRepo repositories.UserRepository
}

func NewGetUsersUseCase(userRepo repositories.UserRepository) *GetUsersUseCase {
	return &GetUsersUseCase{userRepo: userRepo}
}

func (uc *GetUsersUseCase) Execute(ctx context.Context) ([]GetUsersResponse, error) {
	users, err := uc.userRepo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var responses []GetUsersResponse
	for _, user := range users {
		responses = append(responses, GetUsersResponse{
			ID: user.ID.String(),
			Name: user.Name,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		})
	}

	return responses, nil
}