package model

import "time"

type User struct {
	ID        string `gorm:"primaryKey;default:uuid_generate_v4()"`
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:null"`

	UserDetail UserDetail
}
