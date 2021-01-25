package entity

import "time"

type Progress struct {
	BaseEntity
	RoutineID uint64
	Date      time.Time
}

func (Progress) TableName() string {
	return "progress"
}
