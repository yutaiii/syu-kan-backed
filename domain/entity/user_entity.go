package entity

type User struct {
	BaseEntity
	Name        string
	FirebaseUid string
}
