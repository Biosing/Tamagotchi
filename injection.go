package main

import (
	"Tamagotchi/internal/common/middleware"
	"Tamagotchi/internal/database"
	"Tamagotchi/internal/handler"
	"Tamagotchi/internal/repository"
	"Tamagotchi/internal/service"
	"log"

	"Tamagotchi/internal/handler/http"

	"github.com/gin-gonic/gin"
)

func inject(d *database.DataSources) (*gin.Engine, error) {
	log.Println("Start load inject dependencies...")

	authRepo := repository.NewAuthRepository(d.DB)
	authService := service.NewAuthService(authRepo)

	authMiddleware := middleware.NewAuthMiddleware(authRepo)

	service := handler.Services{
		Auth:           authService,
		AuthMiddleware: authMiddleware,
	}

	router := http.SetupRouter(service)

	log.Println("Finish load inject dependencies...")
	return router, nil
}
