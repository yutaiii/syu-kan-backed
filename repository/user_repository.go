package repository

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
	"github.com/yutaiii/syu-kan-backend/domain/model"
	"github.com/yutaiii/syu-kan-backend/store"
	"gorm.io/gorm"
)

type UserRepository struct {
	ctx   context.Context
	store *store.UserStore
}

func NewUserRepository(ctx context.Context) *UserRepository {
	return &UserRepository{
		ctx:   ctx,
		store: store.NewUserStore(ctx),
	}
}

func (r *UserRepository) CreateUser(db *gorm.DB, model *model.User) error {
	entity := r.convertModelToEntity(model)
	return r.store.CreateUser(db, entity)
}

func (r *UserRepository) convertModelToEntity(m *model.User) *entity.User {
	return &entity.User{
		Name:         m.Name,
		FireabaseUid: m.FireabaseUid,
	}
}
