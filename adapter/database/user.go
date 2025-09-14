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
	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FetchWithPosts(ctx context.Context, id int64) (*models.UserWithPosts, error) {
	const userQuery = "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
	var user models.User
	err := r.db.QueryRowContext(ctx, userQuery, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	const postsQuery = "SELECT id, user_id, title, content, created_at, updated_at FROM posts WHERE user_id = $1"
	rows, err := r.db.QueryContext(ctx, postsQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.UserWithPosts{User: user, Posts: []models.Post{}}, nil
		}
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				continue
			}
			return nil, err
		}
		posts = append(posts, p)
	}

	return &models.UserWithPosts{User: user, Posts: posts}, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, user *models.User) (int64, error) {
	const query = "INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id"
	var id int64
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *userRepositoryImpl) Update(ctx context.Context, user *models.User) error {
	const query = "UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.ID)
	return err
}

func (r *userRepositoryImpl) Delete(ctx context.Context, id int64) error {
	const query = "DELETE FROM users WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
