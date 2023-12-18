package routes

import (
	"github.com/gin-gonic/gin"
	controllers "monkey-in-mountain-pass/controllers"
)

func Login(ctx *gin.RouterGroup) {
	ctx.GET("/line/authorize", controllers.HandleLineLoginCallback)
}
