package repositories

import (
	"Hello_World/myapp/internal/domain/entities"
	"github.com/google/uuid"
	"context"
)
type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindAll(ctx context.Context) ([]*entities.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}