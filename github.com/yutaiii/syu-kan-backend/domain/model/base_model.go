package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID          uint64         `json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletededAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}
