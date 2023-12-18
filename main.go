package main

import (
	routes "monkey-in-mountain-pass/middlewares/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
)

func main() {
	Server := gin.Default()
	Server.LoadHTMLGlob("templates/*.html")
	Server.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "test123"})
	})

	api := Server.Group("/api")
	routes.User(api.Group("users"))
	routes.Login(api.Group("login"))
	Server.Run()
}
