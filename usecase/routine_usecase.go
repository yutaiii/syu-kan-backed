package usecase

import (
	"context"
	"log"

	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/repository"
	"github.com/yutaiii/syu-kan-backend/tool/util"
)

type RoutineUsecase struct {
	ctx        context.Context
	repository *repository.RoutineRepository
}

func NewRoutineUsecase(ctx context.Context) *RoutineUsecase {
	return &RoutineUsecase{
		ctx:        ctx,
		repository: repository.NewRoutineRepository(ctx),
	}
}

func (u *RoutineUsecase) GetAllRoutines() ([]*model.Routine, error) {
	db := util.GetConn()
	routines, err := u.repository.GetAllRoutines(db)
	if err != nil {
		log.Printf("RoutineUsecase, GetAllRoutines error: %+v", err)
		return nil, err
	}
	return routines, nil
}

func (u *RoutineUsecase) CreateRoutines(models []*model.Routine) ([]*model.Routine, error) {
	db := util.GetConn()
	routines, err := u.repository.CreateRoutines(models, db)
	if err != nil {
		log.Printf("RoutineUsecase, CreateRoutines error: %+v", err)
		return nil, err
	}
	return routines, nil
}
