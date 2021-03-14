package repository

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/store"

	"gorm.io/gorm"
)

type ProgressRepository struct {
	ctx   context.Context
	store *store.ProgressStore
}

func NewProgressRepository(ctx context.Context) *ProgressRepository {
	return &ProgressRepository{
		ctx:   ctx,
		store: store.NewProgressStore(ctx),
	}
}

func (r *ProgressRepository) GetProgressOfToday(db *gorm.DB, models []*model.Routine) ([]*model.Progress, error) {
	var argRoutineIds []uint64
	query := "`routine_id` IN ? AND `date` >= CURDATE() AND `date` < CURDATE() + INTERVAL 1 DAY"

	for i := 0; i < len(models); i++ {
		argRoutineIds = append(argRoutineIds, models[i].ID)
	}

	// errorになる argsが配列内配列になっている
	//progress, err := r.store.GetByConditions(db, query, argRoutineIds)
	progress, err := r.store.GetByRoutineIds(db, query, argRoutineIds)
	if err != nil {
		return nil, err
	}

	m := model.NewMultiProgress(progress)
	return m, nil
}

func (r *ProgressRepository) CreateProgress(db *gorm.DB, models []*model.Progress) error {
	entities := r.convertModelToEntity(models)
	// 更新するものがない場合は処理を終了
	if len(entities) < 1 {
		return nil
	}
	return r.store.CreateProgress(db, entities)
}

func (r *ProgressRepository) convertModelToEntity(models []*model.Progress) []*entity.Progress {
	var entities []*entity.Progress
	for _, m := range models {
		// 達成していないときはskip
		if !m.Achieved {
			continue
		}
		e := &entity.Progress{
			RoutineID: m.RoutineID,
			Date:      m.Date,
		}
		entities = append(entities, e)
	}
	return entities
}
