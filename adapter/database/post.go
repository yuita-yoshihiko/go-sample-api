package database

import (
	"context"
	"database/sql"

	"github.com/yuita-yoshihiko/go-sample-api/infrastructure/db"
	"github.com/yuita-yoshihiko/go-sample-api/models"
	"github.com/yuita-yoshihiko/go-sample-api/usecase/repository"
)

type PostRepositoryImpl struct {
	db db.DBUtils
}

func NewPostRepository(db db.DBUtils) repository.PostRepository {
	return &PostRepositoryImpl{db: db}
}

func (r *PostRepositoryImpl) FetchByUserID(ctx context.Context, userID int64) ([]*models.Post, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, user_id, title, content, created_at, updated_at FROM posts WHERE user_id = $1", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var post []*models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
		}
		post = append(post, &p)
	}
	return post, nil
}
