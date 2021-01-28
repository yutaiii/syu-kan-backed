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

func GetAllRoutines() ([]*model.Routine, error) {
	db := util.GetConn()
	routines, err := repository.GetAllRoutines(db)
	if err != nil {
		log.Printf("GetAllRoutines err: %+v", err)
		return nil, err
	}
	return routines, nil
}

func (u *RoutineUsecase) CreateRoutines(models []*model.Routine) ([]*model.Routine, error) {
	db := util.GetConn()
	routines, err := u.repository.CreateRoutines(models, db)
	if err != nil {
		log.Printf("RoutineUsecase, CreateRoutines err: %+v", err)
		return nil, err
	}
	return routines, nil
}
