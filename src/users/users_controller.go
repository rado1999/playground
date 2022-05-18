package users

import (
	"github.com/gin-gonic/gin"
)

func UsersController(server *gin.Engine) {
	server.GET("/users", GetUsers)
	server.POST("/users", CreateUser)
}
