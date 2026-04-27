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

func (r *MySQLUserRepository) Update(ctx context.Context, user *entities.User) error {
	query := `UPDATE users SET name = ?, email = ?, updated_at = ? WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.UpdatedAt,
		user.ID.String(),
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return apperror.NewResourceNotFoundError("user not found")
	}

	return nil
}

func (r *MySQLUserRepository) ChangePassword(ctx context.Context, id uuid.UUID, newPassword string) error {
	query := `UPDATE users SET password = ? WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, newPassword, id.String())
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return apperror.NewResourceNotFoundError("user not found")
	}

	return nil
}

func (r *MySQLUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id.String())
	if err != nil{
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return apperror.NewResourceNotFoundError("user not found")
	}

	return nil
}

func (r *MySQLUserRepository) FindAll(ctx context.Context) ([]*entities.User, error) {
	query := `SELECT id, name, email, password, created_at, updated_at FROM users`

	rows, err := r.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
			return nil, err
		}	

	return users, nil
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
		return nil, apperror.NewResourceNotFoundError("user not found")
	}
	if err != nil {
		return nil, err
	}

	user.ID, err = uuid.Parse(rawID)
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
	if err != nil {
		return nil, err
	}

	user.ID, _ = uuid.Parse(rawID) 

	return user, nil
}