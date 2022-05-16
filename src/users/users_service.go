package users

import "github.com/gin-gonic/gin"

func GetUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "there is no users"})
}
