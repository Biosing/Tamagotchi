package main

import (
	"Tamagotchi/internal/handler"

	"github.com/gin-gonic/gin"

	_ "Tamagotchi/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger embed files

func SetupRouter(s handler.Services) *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	{
		userGroup := apiV1.Group("/user")
		{
			userGroup.GET("/", s.User.GetUsers)
			userGroup.POST("/", s.User.CreateUser)
			userGroup.GET("/:id", s.User.GetUser)
			userGroup.PUT("/:id", s.User.UpdateUser)
			userGroup.DELETE("/:id", s.User.DeleteUser)
		}
	}
	// Маршруты для пользователей

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
