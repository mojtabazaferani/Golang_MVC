package routes

import "github.com/gin-gonic/gin"

func Route() *gin.Engine{

	Route := gin.Default()

	Route.GET("/", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{"message": "Hello World"})

	})

	return Route

}