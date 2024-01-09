package main

import (
	"Tamagotchi/handlers"

	"github.com/gin-gonic/gin"

	_ "Tamagotchi/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger embed files

func SetupRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		userGroup := apiV1.Group("/user")
		{
			userGroup.GET("/", handlers.GetUsers)
			userGroup.POST("/", handlers.CreateUser)
			userGroup.GET("/:id", handlers.GetUser)
			userGroup.PUT("/:id", handlers.UpdateUser)
			userGroup.DELETE("/:id", handlers.DeleteUser)
		}
	}
	// Маршруты для пользователей

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
