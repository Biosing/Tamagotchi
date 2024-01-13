package http

import (
	"Tamagotchi/internal/handler"

	"github.com/gin-gonic/gin"

	_ "Tamagotchi/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(s handler.Services) *gin.Engine {
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	{
		SetupAuthenticationRouter(apiV1, s)
	}
	// Маршруты для пользователей

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
