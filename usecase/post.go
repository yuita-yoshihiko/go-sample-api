package usecase

import (
	"context"

	"github.com/yuita-yoshihiko/go-sample-api/models"
	"github.com/yuita-yoshihiko/go-sample-api/usecase/repository"
)

type PostUseCase interface {
	FetchByUserID(context.Context, int64) ([]*models.Post, error)
	FetchByUserIDWithComments(context.Context, int64) ([]*models.PostWithComments, error)
}

type postUseCaseImpl struct {
	repository repository.PostRepository
}

func NewPostUseCase(
	r repository.PostRepository,
) PostUseCase {
	return &postUseCaseImpl{
		repository: r,
	}
}

func (p *postUseCaseImpl) FetchByUserID(ctx context.Context, userID int64) ([]*models.Post, error) {
	return p.repository.FetchByUserID(ctx, userID)
}

func (p *postUseCaseImpl) FetchByUserIDWithComments(ctx context.Context, userID int64) ([]*models.PostWithComments, error) {
	return p.repository.FetchByUserIDWithComments(ctx, userID)
}
