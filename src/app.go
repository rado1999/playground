package src

import (
	"project/config"
	"project/src/users"

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

	return server
}

func setControllers(server *gin.Engine) {
	users.UsersController(server)
}
