package database

import (
	"context"
	"database/sql"

	"github.com/yuita-yoshihiko/go-sample-api/infrastructure/db"
	"github.com/yuita-yoshihiko/go-sample-api/models"
	"github.com/yuita-yoshihiko/go-sample-api/usecase/repository"
)

type userRepositoryImpl struct {
	db db.DBUtils
}

func NewUserRepository(db db.DBUtils) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Fetch(ctx context.Context, id int64) (*models.User, error) {
	const query = "SELECT * FROM users WHERE id = $1"
	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *models.User) error {
	const query = "INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at"
	return r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *userRepositoryImpl) Update(ctx context.Context, user *models.User) error {
	const query = "UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3 RETURNING updated_at"
	return r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.ID).Scan(&user.UpdatedAt)
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id int64) error {
	const query = "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
