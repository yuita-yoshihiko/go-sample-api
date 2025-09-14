package converter

import (
	"github.com/yuita-yoshihiko/go-sample-api/models"
)

type UserOutput struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

type UserConverter interface {
	ToUserOutput(*models.User) *UserOutput
}

type userConverterImpl struct {
}

func NewUserConverter() UserConverter {
	return &userConverterImpl{}
}

func (p *userConverterImpl) ToUserOutput(input *models.User) *UserOutput {
	return &UserOutput{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
	}
}
