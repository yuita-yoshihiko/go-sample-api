package repository

import (
	"context"

	"github.com/yuita-yoshihiko/go-sample-api/models"
)

type UserRepository interface {
	Fetch(ctx context.Context, id int64) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id int64) error
}
