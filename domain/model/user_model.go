package model

import "github.com/yutaiii/syu-kan-backend/domain/entity"

type InputUser struct {
	BaseModel
	Name        string `json:"name"`
	FirebaseUid string `json:"firebaseUid"`
}

type OutputUser struct {
	BaseModel
	Name        string `json:"name"`
	FirebaseUid string `json:"-"`
}

func NewOutputUser(entity *entity.User) *OutputUser {
	return &OutputUser{
		BaseModel: BaseModel{
			ID:          entity.BaseEntity.ID,
			CreatedAt:   entity.BaseEntity.CreatedAt,
			UpdatedAt:   entity.BaseEntity.UpdatedAt,
			DeletededAt: entity.BaseEntity.DeletedAt,
		},
		Name: entity.Name,
	}
}
