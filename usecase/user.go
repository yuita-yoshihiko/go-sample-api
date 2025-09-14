package usecase

import (
	"context"

	"github.com/yuita-yoshihiko/go-sample-api/models"
	"github.com/yuita-yoshihiko/go-sample-api/usecase/repository"
)

type UserUseCase interface {
	Fetch(context.Context, int64) (*models.User, error)
	FetchWithPosts(context.Context, int64) (*models.UserWithPosts, error)
	Create(context.Context, *models.User) error
	Update(context.Context, *models.User) error
	Delete(context.Context, int64) error
}

type userUseCaseImpl struct {
	repository repository.UserRepository
}

func NewUserUseCase(
	r repository.UserRepository,
) UserUseCase {
	return &userUseCaseImpl{
		repository: r,
	}
}

func (u *userUseCaseImpl) Fetch(ctx context.Context, id int64) (*models.User, error) {
	return u.repository.Fetch(ctx, id)
}

func (u *userUseCaseImpl) FetchWithPosts(ctx context.Context, id int64) (*models.UserWithPosts, error) {
	return u.repository.FetchWithPosts(ctx, id)
}

func (u *userUseCaseImpl) Create(ctx context.Context, user *models.User) error {
	return u.repository.Create(ctx, user)
}

func (u *userUseCaseImpl) Update(ctx context.Context, user *models.User) error {
	return u.repository.Update(ctx, user)
}

func (u *userUseCaseImpl) Delete(ctx context.Context, id int64) error {
	return u.repository.Delete(ctx, id)
}
