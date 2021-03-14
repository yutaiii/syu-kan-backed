package model

import (
	"time"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
)

type Progress struct {
	BaseModel
	RoutineID uint64    `json:"routineId"`
	Date      time.Time `json:"date"`
	Achieved  bool      `json:"achieved"`
}

func NewMultiProgress(entities []*entity.Progress) []*Progress {
	var models []*Progress
	for _, e := range entities {
		model := NewProgress(e)
		models = append(models, model)
	}
	return models
}

func NewProgress(entity *entity.Progress) *Progress {
	return &Progress{
		BaseModel: BaseModel{
			ID:          entity.BaseEntity.ID,
			CreatedAt:   entity.BaseEntity.CreatedAt,
			UpdatedAt:   entity.BaseEntity.UpdatedAt,
			DeletededAt: entity.BaseEntity.DeletedAt,
		},
		RoutineID: entity.RoutineID,
		Date:      entity.Date,
	}
}
