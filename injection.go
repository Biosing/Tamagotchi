package main

import (
	"Tamagotchi/internal/database"
	"Tamagotchi/internal/handler"
	"Tamagotchi/internal/repository"
	"Tamagotchi/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func inject(d *database.DataSources) (*gin.Engine, error) {
	log.Println("Start load inject dependencies...")

	userRepo := repository.NewUserRepository(d.DB)

	userService := service.NewUserService(userRepo)

	service := handler.Services{
		User: userService,
	}

	router := SetupRouter(service)

	log.Println("Finish load inject dependencies...")
	return router, nil
}
