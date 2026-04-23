package repositories

import (
	"Hello_World/myapp/internal/domain/entities"
	"Hello_World/myapp/pkg/apperror"
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	return err
}

func (r * MySQLUserRepository) FindById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?`

	var rawID string
	user:= &entities.User{}

	err := r.db.QueryRowContext(ctx, query, id.String()).Scan(
		&rawID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, apperror.NewResourceNotFoundError("User not found")
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MySQLUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?`

	var rawID string
	user := &entities.User{}

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&rawID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} 
	if (err != nil) {
		return nil, err
	}

	user.ID, _ = uuid.Parse(rawID) 

	return user, nil
}

func (r *MySQLUserRepository) Update(ctx context.Context, user *entities.User) error {
	query := `UPDATE users SET name = ?, email = ?, password = ?, updated_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
		user.UpdatedAt,
		user.ID.String(),
	)
	return err
}

func (r *MySQLUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id.String())
	return err
}