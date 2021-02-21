package model

import "github.com/yutaiii/syu-kan-backend/domain/entity"

type User struct {
	BaseModel
	Name        string `json:"name"`
	FirebaseUid string `json:"firebaseUid"`
}

func NewUser(entity *entity.User) *User {
	return &User{
		BaseModel: BaseModel{
			ID:          entity.BaseEntity.ID,
			CreatedAt:   entity.BaseEntity.CreatedAt,
			UpdatedAt:   entity.BaseEntity.UpdatedAt,
			DeletededAt: entity.BaseEntity.DeletedAt,
		},
		Name:        entity.Name,
		FirebaseUid: entity.FirebaseUid,
	}
}
