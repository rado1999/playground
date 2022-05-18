package models

import (
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, id uint) User {
	var user User
	db.Find(&user, id)
	return user
}

func CreateUser(db *gorm.DB, user UserDto) {
	db.Create(User{
		Username: user.Username,
		Age:      user.Age,
	})
}
