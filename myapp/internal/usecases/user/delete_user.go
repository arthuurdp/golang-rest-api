package user

import (
	"Hello_World/myapp/internal/domain/repositories"
	"context"

	"github.com/google/uuid"
)

type DeleteUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewDeleteUserUseCase(userRepo repositories.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{userRepo: userRepo}
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context,id uuid.UUID) error {
	return uc.userRepo.Delete(ctx, id)
}