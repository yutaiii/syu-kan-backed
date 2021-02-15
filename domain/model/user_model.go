package model

type User struct {
	BaseModel
	Name         string `json:"name"`
	FireabaseUid string `json:"firebaseUid"`
}
