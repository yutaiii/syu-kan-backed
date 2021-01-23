package model

import (
	"time"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
)

type Routine struct {
	BaseModel
	Name      string    `json:"name"`
	StartedAt time.Time `json:"startedAt"`
}

func NewRoutine(entity *entity.Routine) *Routine {
	return &Routine{
		BaseModel: BaseModel{
			ID:          entity.BaseEntity.ID,
			CreatedAt:   entity.BaseEntity.CreatedAt,
			UpdatedAt:   entity.BaseEntity.UpdatedAt,
			DeletededAt: entity.BaseEntity.DeletedAt,
		},
		Name:      entity.Name,
		StartedAt: entity.StartedAt,
	}
}

func NewRoutines(entities []*entity.Routine) []*Routine {
	var models []*Routine
	for _, e := range entities {
		model := NewRoutine(e)
		models = append(models, model)
	}
	return models
}
