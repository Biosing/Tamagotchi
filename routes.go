package main

import (
	"Tamagotchi/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Маршруты для пользователей
	userGroup := r.Group("/user")
	{
		userGroup.GET("/", handlers.GetUsers)
		userGroup.POST("/", handlers.CreateUser)
		userGroup.GET("/:id", handlers.GetUser)
		userGroup.PUT("/:id", handlers.UpdateUser)
		userGroup.DELETE("/:id", handlers.DeleteUser)
	}

	return r
}
