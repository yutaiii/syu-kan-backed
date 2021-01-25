package repository

import (
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/store"

	"gorm.io/gorm"
)

func GetAllRoutines(db *gorm.DB) ([]*model.Routine, error) {
	routines, err := store.GetAll(db)
	if err != nil {
		return nil, err
	}
	models := model.NewRoutines(routines)
	return models, err
}
