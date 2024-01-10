package model

import (
	"context"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetUsers(ctx context.Context)
	CreateUser(ctx context.Context)
	GetUser(ctx context.Context)
	UpdateUser(ctx context.Context)
	DeleteUser(ctx context.Context)
}

type UserService interface {
	GetUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
