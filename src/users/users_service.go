package users

import (
	"net/http"
	dbconnection "project/db_connection"
	"project/src/users/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(ctx *gin.Context) {
	var userDto models.UserDto
	var user models.User
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbconnection.DB.Find(&user, models.User{Username: userDto.Username})
	ctx.JSON(200, user)
}

func CreateUser(ctx *gin.Context) {
	var userDto models.UserDto
	if err := ctx.BindJSON(&userDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: userDto.Username,
		Age:      userDto.Age,
	}

	dbconnection.DB.Create(&user)
}
