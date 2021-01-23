package store

import (
	"github.com/jinzhu/gorm"
	"github.com/yutaiii/command-box-backend/entity"
)

func GetAll(db *gorm.DB) ([]*entity.Routine, error) {
	// routines := make([]*.Routine, 0)
	routines := make([]*model.Routine, 0)
	//routines := &[]entity.Routine{}
	err := db.Find(routines)
	if err != nil {
		return nil, err
	}
	return routines, nil
}
