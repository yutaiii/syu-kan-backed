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

func (s *ProgressStore) CreateProgress(db *gorm.DB, entities []*entity.Progress) error {
	return db.Create(&entities).Error
}
