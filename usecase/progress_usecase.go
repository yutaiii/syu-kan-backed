package usecase

import (
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/repository"
	"github.com/yutaiii/syu-kan-backend/tool/util"
)

func CreateProgress(models []*model.Progress) error {
	db := util.GetConn()
	return repository.CreateProgress(db, models)
}
