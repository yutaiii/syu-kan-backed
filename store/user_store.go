package store

import (
	"context"
	"fmt"

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

func (s *UserStore) FindUserByCondition(db *gorm.DB, field, value string) (*entity.User, error) {
	u := &entity.User{}
	query := fmt.Sprintf("%s = ?", field)
	err := db.Where(query, value).First(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
