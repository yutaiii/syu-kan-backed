package store

import (
	"context"

	"github.com/yutaiii/syu-kan-backend/domain/entity"
	"gorm.io/gorm"
)

type UserStore struct {
	ctx context.Context
}

func NewUserStore(ctx context.Context) *UserStore {
	return &UserStore{
		ctx: ctx,
	}
}

func (s *UserStore) CreateUser(db *gorm.DB, user *entity.User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
