package src

import (
	"project/config"
	dbconnection "project/db_connection"
	"project/src/users"
	"project/src/users/models"

	"github.com/gin-gonic/gin"
)

var conf = config.New()

func App() *gin.Engine {
	server := gin.New()

	if conf.GIN_MODE != "release" {
		server.Use(gin.Logger(), gin.Recovery())
	}

	server.SetTrustedProxies(nil)

	setControllers(server)
	setModels()

	return server
}

func setControllers(server *gin.Engine) {
	users.UsersController(server)
}

func setModels() {
	dbconnection.DB.AutoMigrate(&models.User{})
}
