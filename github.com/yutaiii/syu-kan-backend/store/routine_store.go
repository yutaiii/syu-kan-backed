package store

import (
	"github.com/yutaiii/syu-kan-backend/domain/entity"
	"gorm.io/gorm"
)

func GetAll(db *gorm.DB) ([]*entity.Routine, error) {
	routines := make([]*entity.Routine, 0)
	err := db.Find(&routines).Error
	if err != nil {
		return nil, err
	}
	return routines, nil
}
