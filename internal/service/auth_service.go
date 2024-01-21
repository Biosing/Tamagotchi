package service

import (
	"Tamagotchi/internal/model"
	"Tamagotchi/internal/schema"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authRepository model.AuthRepository
}

func NewAuthService(repo model.AuthRepository) model.AuthService {
	return &authService{
		authRepository: repo,
	}
}

func (nas *authService) Registration(ctx context.Context, payload *schema.RegistrationSchema) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a new user with the hashed password
	user := &model.User{
		Login:     payload.Login,
		Email:     payload.Email,
		Password:  string(hashedPassword),
		FullName:  payload.FullName,
		CreatedAt: time.Now(),
	}

	// Save the user to the repository
	err = nas.authRepository.Registration(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (nas *authService) Authorization(ctx *gin.Context, payload *schema.AuthorizationSchema) error {

	user, err := nas.authRepository.GetUserByEmailOrLogin(ctx, payload.LoginOrEmail)
	if err != nil {
		return err
	}

	// Compare the passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return err
	}

	sessionId := uuid.New().String()
	expirationHours := viper.GetString("COOKIE_Expiration_Hours")
	cookName := viper.GetString("COOKIE_NAME")
	hours, err := strconv.Atoi(expirationHours)
	if err != nil {
		return err
	}
	expiresAt := time.Now().Add(time.Hour * time.Duration(hours)).UTC()

	// Create a cookie
	cookie := &http.Cookie{
		Name:     cookName,
		Value:    sessionId,
		Expires:  expiresAt,
		HttpOnly: true,
	}

	// Set the cookie in the response
	http.SetCookie(ctx.Writer, cookie)

	nas.authRepository.SaveSession(ctx, &model.Session{
		Id:        sessionId,
		UserId:    user.Id,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
		UserAgent: ctx.Request.UserAgent(),
		IPAddress: ctx.ClientIP(),
	})

	return nil
}

func (nas *authService) Logout(ctx *gin.Context) error {

	cookName := viper.GetString("COOKIE_NAME")

	cookie, err := ctx.Request.Cookie(cookName)
	if err != nil {
		return err
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(ctx.Writer, cookie)

	return nil
}
