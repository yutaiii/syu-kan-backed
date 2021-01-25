package entity

import (
	"time"
)

type Routine struct {
	BaseEntity
	Name      string
	StartedAt time.Time
}
