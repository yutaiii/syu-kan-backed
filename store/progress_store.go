package store

import (
	"github.com/yutaiii/syu-kan-backend/domain/entity"

	"gorm.io/gorm"
)

func CreateProgress(db *gorm.DB, entities []*entity.Progress) error {
	return db.Create(&entities).Error
}
