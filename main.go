package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tonyseptiann12/go-crud/controllers"
	"github.com/tonyseptiann12/go-crud/initializer"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.DELETE("/posts/:id", controllers.PostsDelete)
	router.Run()
}
