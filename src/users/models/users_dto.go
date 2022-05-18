package models

type UserDto struct {
	Username string `json:"username" binding:"required"`
	Age      uint   `json:"age" binding:"required"`
}
