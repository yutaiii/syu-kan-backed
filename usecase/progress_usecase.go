package usecase

import (
	"context"
	"log"

	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/repository"
	"github.com/yutaiii/syu-kan-backend/tool/util"
)

type ProgressUsecase struct {
	ctx                context.Context
	progressRepository *repository.ProgressRepository
	routineRepository  *repository.RoutineRepository
}

func NewProgressUsecase(ctx context.Context) *ProgressUsecase {
	return &ProgressUsecase{
		ctx:                ctx,
		progressRepository: repository.NewProgressRepository(ctx),
		routineRepository:  repository.NewRoutineRepository(ctx),
	}
}

func (u *ProgressUsecase) GetProgressOfToday(m *model.RoutineForGetInput) ([]*model.Progress, error) {
	db := util.GetConn()
	routines, err := u.routineRepository.GetByUserId(db, m)
	if err != nil {
		log.Printf("RoutineUsecase, GetProgressOfToday, GetByUserId error: %+v", err)
		return nil, err
	}

	// ルーティーンが登録されていない場合
	if len(routines) < 1 {
		log.Print("RoutineUsecase, GetProgressOfToday, Routine is not registered")
		var emptyModel []*model.Progress
		return emptyModel, nil
	}
	progress, err := u.progressRepository.GetProgressOfToday(db, routines)
	if err != nil {
		log.Printf("RoutineUsecase, GetProgressOfToday, GetProgressOfToday error: %+v", err)
		return nil, err
	}
	return progress, nil
}

func (u *ProgressUsecase) CreateProgress(models []*model.Progress) error {
	db := util.GetConn()
	return u.progressRepository.CreateProgress(db, models)
}
