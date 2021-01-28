package repository

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/store"

	"gorm.io/gorm"
)

type RoutineRepository struct {
	ctx   context.Context
	store *store.RoutineStore
}

func NewRoutineRepository(ctx context.Context) *RoutineRepository {
	return &RoutineRepository{
		ctx:   ctx,
		store: store.NewRoutineStore(ctx),
	}
}

func GetAllRoutines(db *gorm.DB) ([]*model.Routine, error) {
	routines, err := store.GetAll(db)
	if err != nil {
		return nil, err
	}
	models := model.NewRoutines(routines)
	return models, err
}

func (r *RoutineRepository) CreateRoutines(models []*model.Routine, db *gorm.DB) ([]*model.Routine, error) {
	entities := r.convertModelsToEntity(models)
	result, err := r.store.CreateRoutines(entities, db)
	if err != nil {
		return nil, err
	}
	m := model.NewRoutines(result)
	return m, nil
}

func (r *RoutineRepository) convertModelsToEntity(models []*model.Routine) []*entity.Routine {
	var entities []*entity.Routine
	for _, m := range models {
		e := r.convertModelToEntity(m)
		entities = append(entities, e)
	}
	return entities
}

func (r *RoutineRepository) convertModelToEntity(m *model.Routine) *entity.Routine {
	return &entity.Routine{
		Name:      m.Name,
		StartedAt: m.StartedAt,
	}
}
