package model

import "time"

type Progress struct {
	BaseModel
	RoutineID uint64    `json:"routineId"`
	Date      time.Time `json:"date"`
	Achieved  bool      `json:"achieved"`
}
