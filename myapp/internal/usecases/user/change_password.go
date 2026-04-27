package user

import (
	"Hello_World/myapp/internal/domain/repositories"
	"Hello_World/myapp/pkg/apperror"
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordRequest struct {
	Password string 
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
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err == nil {
		return apperror.NewConflictError("new password can't be the same as the current one")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return apperror.NewInternalServerError("error hashing password")
	} 

	return uc.userRepo.ChangePassword(ctx, id, string(hash))
} 