package usecase

import (
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/repository"
	"github.com/yutaiii/syu-kan-backend/tool/util"
)

func GetAllRoutines() ([]*model.Routine, error) {
	db := util.GetConn()
	return repository.GetAllRoutines(db)
}
