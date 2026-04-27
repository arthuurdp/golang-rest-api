package user

import (
	"Hello_World/myapp/internal/domain/repositories"
	"Hello_World/myapp/pkg/apperror"
	"context"

	"github.com/google/uuid"
)

type ChangePasswordRequest struct {
	Password string `json:"password" binding:"required,min=6"`
}

type ChangePasswordUseCase struct {
	userRepo repositories.UserRepository
}

func NewChangePasswordUseCase(userRepo repositories.UserRepository) *ChangePasswordUseCase {
	return &ChangePasswordUseCase{ userRepo: userRepo}
}

func (uc *ChangePasswordUseCase) Execute(ctx context.Context, id uuid.UUID, req ChangePasswordRequest) error {
	user, err := uc.userRepo.FindById(ctx, id)
	if err != nil {
		return apperror.NewResourceNotFoundError("user not found")
	}

	if req.Password == user.Password {
		return apperror.NewConflictError("new password can't be the same as the current one")
	}

	return uc.userRepo.ChangePassword(ctx, id, req.Password)
} 