package usecase

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/repository"
	"github.com/yutaiii/syu-kan-backend/tool/util"
)

type ProgressUsecase struct {
	ctx        context.Context
	repository *repository.ProgressRepository
}

func NewProgressUsecase(ctx context.Context) *ProgressUsecase {
	return &ProgressUsecase{
		ctx:        ctx,
		repository: repository.NewProgressRepository(ctx),
	}
}

func (u *ProgressUsecase) CreateProgress(models []*model.Progress) error {
	db := util.GetConn()
	return u.repository.CreateProgress(db, models)
}
