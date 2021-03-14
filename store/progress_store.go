package store

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/entity"

	"gorm.io/gorm"
)

type ProgressStore struct {
	ctx context.Context
}

func NewProgressStore(ctx context.Context) *ProgressStore {
	return &ProgressStore{
		ctx: ctx,
	}
}

func (s *ProgressStore) GetByConditions(db *gorm.DB, query string, args ...interface{}) ([]*entity.Progress, error) {
	progress := make([]*entity.Progress, 0)
	err := db.Where(query, args).Find(&progress).Error
	if err != nil {
		return nil, err
	}
	return progress, nil
}

func (s *ProgressStore) GetByRoutineIds(db *gorm.DB, query string, routineIds []uint64) ([]*entity.Progress, error) {
	progress := make([]*entity.Progress, 0)
	err := db.Where(query, routineIds).Find(&progress).Error
	if err != nil {
		return nil, err
	}
	return progress, nil
}

func (s *ProgressStore) CreateProgress(db *gorm.DB, entities []*entity.Progress) error {
	return db.Create(&entities).Error
}
