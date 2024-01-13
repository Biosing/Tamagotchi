package model

import (
	"Tamagotchi/internal/schema"
	"context"

	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	Registration(ctx context.Context, user *User) error
	GetUserByEmailOrLogin(ctx context.Context, loginOrEmail string) (*User, error)
	SaveSession(c context.Context, session *Session) error
	CheckSession(c context.Context, sessionID string) (*Session, error)
}

type AuthService interface {
	Registration(ctx context.Context, payload *schema.RegistrationSchema) error
	Authorization(ctx *gin.Context, payload *schema.AuthorizationSchema) error
	Logout(ctx *gin.Context) error
}
