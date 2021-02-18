package entity

import (
	"time"
)

type Routine struct {
	BaseEntity
	Name      string
	UserID    uint64
	StartedAt time.Time
}
