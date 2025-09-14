package repository

import (
	"context"

	"github.com/yuita-yoshihiko/go-sample-api/models"
)

type PostRepository interface {
	FetchByUserID(ctx context.Context, userID int64) ([]*models.Post, error)
	FetchByUserIDWithComments(ctx context.Context, userID int64) ([]*models.PostWithComments, error)
}
