package usecase

import (
	"context"
	"github.com/yuita-yoshihiko/go-sample-api/infrastructure/db"
	"github.com/yuita-yoshihiko/go-sample-api/usecase/converter"
	"github.com/yuita-yoshihiko/go-sample-api/usecase/repository"
)

type UserUseCase interface {
	Fetch(context.Context, int64) (*converter.UserOutput, error)
}

type userUseCaseImpl struct {
	dbUtils    db.DBUtils
	repository repository.UserRepository
	converter  converter.UserConverter
}

func NewUserUseCase(
	du db.DBUtils,
	r repository.UserRepository,
	c converter.UserConverter,
) UserUseCase {
	return &userUseCaseImpl{
		dbUtils:    du,
		repository: r,
		converter:  c,
	}
}

func (u *userUseCaseImpl) Fetch(ctx context.Context, id int64) (*converter.UserOutput, error) {
	m, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.ToUserOutput(m), nil
}
