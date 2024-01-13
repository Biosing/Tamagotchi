package middleware

import (
	"Tamagotchi/internal/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type AuthMiddleware struct {
	authRepository model.AuthRepository
}

func NewAuthMiddleware(authRepository model.AuthRepository) *AuthMiddleware {
	return &AuthMiddleware{
		authRepository: authRepository,
	}
}

func (am *AuthMiddleware) CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		cookName := viper.GetString("COOKIE_NAME")

		cookie, err := c.Request.Cookie(cookName)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Необходима авторизация"})
			return
		}

		if !am.isValidCookie(c.Request.Context(), cookie.Value) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Недействительная сессия"})
			return
		}

		c.Next()
	}
}

func (am *AuthMiddleware) isValidCookie(ctx context.Context, cookieValue string) bool {
	sessionValid, err := am.authRepository.CheckSession(ctx, cookieValue)

	if err != nil {
		return false
	}

	if sessionValid == nil {
		return false
	}

	if sessionValid.ExpiresAt.Before(time.Now()) {
		return false
	}

	return true
}
