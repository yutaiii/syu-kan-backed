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

// curl -X POST http://localhost:8000/progress -d 'routineId=1' 'date=2020/01/24' 'achieved=true'
// curl -X POST -H "Content-Type: application/json" -d '{"routineId":1, "date":"2020/01/24, "achieved": true}' http://localhost:8000/progress
