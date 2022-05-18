package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Username  string
	Age       uint
	CreatedAt time.Time
}
