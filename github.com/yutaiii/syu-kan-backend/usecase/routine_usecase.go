package usecase

import (
	"log"

	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/repository"
	"github.com/yutaiii/syu-kan-backend/tool/util"
)

func GetAllRoutines() ([]*model.Routine, error) {
	db := util.GetConn()
	routines, err := repository.GetAllRoutines(db)
	if err != nil {
		log.Printf("GetAllRoutines err: %+v", err)
		return nil, err
	}
	return routines, nil
}
