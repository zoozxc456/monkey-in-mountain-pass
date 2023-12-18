package routes

import (
	"github.com/gin-gonic/gin"
	controllers "monkey-in-mountain-pass/controllers"
)

func User(route *gin.RouterGroup) {
	route.GET("/callback",controllers.RegisterByLineLogin)
}
