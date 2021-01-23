package model

import (
	"reflect"
	"time"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
)

type Routine struct {
	BaseModel
	Name      string    `json:"name"`
	StartedAt time.Time `json:"startedAt"`
}

func NewRoutine(entity entity.Routine) *Routine {
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

func NewRoutines(entities *[]entity.Routine) []*Routine {
	var models []*Routine
	e := reflect.TypeOf(entities)
	cnt := e.NumField()
	for i := 0; i < cnt; i++ {
		model := NewRoutine((*entities)[i])
		models = append(models, model)
	}
	return models
}
