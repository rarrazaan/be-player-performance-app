package model

type UserDetail struct {
	ID          string `gorm:"primaryKey;default:uuid_generate_v4()"`
	UserID      string
	FullName    string
	Age         int
	Gender      string
	Address     string
	PhoneNumber string
}
