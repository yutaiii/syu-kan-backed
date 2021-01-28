package store

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
	"gorm.io/gorm"
)

type RoutineStore struct {
	ctx context.Context
}

func NewRoutineStore(ctx context.Context) *RoutineStore {
	return &RoutineStore{
		ctx: ctx,
	}
}

func (s *RoutineStore) GetAll(db *gorm.DB) ([]*entity.Routine, error) {
	routines := make([]*entity.Routine, 0)
	err := db.Find(&routines).Error
	if err != nil {
		return nil, err
	}
	return routines, nil
}

func (s *RoutineStore) CreateRoutines(e []*entity.Routine, db *gorm.DB) ([]*entity.Routine, error) {
	err := db.Create(&e).Error
	if err != nil {
		return nil, err
	}
	return e, nil
}
