package handler

import (
	"Tamagotchi/internal/common/middleware"
	"Tamagotchi/internal/model"
)

type Services struct {
	Auth           model.AuthService
	AuthMiddleware *middleware.AuthMiddleware
}
