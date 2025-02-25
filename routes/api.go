package routes

import (
	"github.com/gin-gonic/gin"
	"myapp/app/Http/Controllers"
)

func Route() *gin.Engine{

	Route := gin.Default()

	Route.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "Hello World"})

	})

	Route.GET("/index", controllers.Index)

	return Route

}