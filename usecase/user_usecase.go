package usecase

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/repository"
	"github.com/yutaiii/syu-kan-backend/tool/util"
)

type UserUsecase struct {
	ctx        context.Context
	repository *repository.UserRepository
}

func NewUserUsecase(ctx context.Context) *UserUsecase {
	return &UserUsecase{
		ctx:        ctx,
		repository: repository.NewUserRepository(ctx),
	}
}

func (u *UserUsecase) CreateUser(model *model.User) error {
	db := util.GetConn()
	return u.repository.CreateUser(db, model)
}
