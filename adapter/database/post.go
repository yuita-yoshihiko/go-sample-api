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

func (r *PostRepositoryImpl) FetchByUserIDWithComments(ctx context.Context, userID int64) ([]*models.PostWithComments, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, user_id, title, content, created_at, updated_at FROM posts WHERE user_id = $1", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var postsWithComments []*models.PostWithComments
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
		}

		commentsRows, err := r.db.QueryContext(ctx, "SELECT id, post_id, user_id, content, created_at, updated_at FROM comments WHERE post_id = $1", p.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				postsWithComments = append(postsWithComments, &models.PostWithComments{Post: p, Comments: []models.Comment{}})
				continue
			}
			return nil, err
		}
		defer commentsRows.Close()

		var comments []models.Comment
		for commentsRows.Next() {
			var c models.Comment
			if err := commentsRows.Scan(&c.ID, &c.PostID, &c.Content, &c.CreatedAt, &c.UpdatedAt); err != nil {
				if err == sql.ErrNoRows {
					break
				}
			}
			comments = append(comments, c)
		}

		postsWithComments = append(postsWithComments, &models.PostWithComments{Post: p, Comments: comments})
	}
	return postsWithComments, nil
}
