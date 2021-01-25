package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
