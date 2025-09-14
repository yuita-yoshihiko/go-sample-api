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
