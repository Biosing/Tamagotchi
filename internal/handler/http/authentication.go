package http

import (
	"Tamagotchi/internal/common/middleware"
	"Tamagotchi/internal/handler"
	"Tamagotchi/internal/schema"

	"github.com/gin-gonic/gin"

	_ "Tamagotchi/docs"
)

type AuthenticationRouter struct {
	Services handler.Services
}

func SetupAuthenticationRouter(apiV1 *gin.RouterGroup, serv handler.Services) *gin.RouterGroup {

	ar := AuthenticationRouter{
		Services: serv,
	}

	userGroup := apiV1.Group("/user")
	{
		userGroup.POST("/registration", middleware.GinErrorHandle(ar.RegistrationHandler))
		userGroup.POST("/auth", middleware.GinErrorHandle(ar.AuthHandler))
		userGroup.GET("/logout", ar.Services.AuthMiddleware.CheckAuth(), middleware.GinErrorHandle(ar.LogoutHandler))
	}

	return userGroup
}

// @Summary Registration
// @Tags users
// @Accept json
// @Produce json
// @Param data body schema.RegistrationSchema true "RegistrationSchema"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /user/registration [post]
func (a *AuthenticationRouter) RegistrationHandler(c *gin.Context) error {
	ctx := c.Request.Context()
	var payload *schema.RegistrationSchema

	if err := c.BindJSON(&payload); err != nil {
		return err
	}

	err := a.Services.Auth.Registration(ctx, payload)
	if err != nil {
		return err
	}
	return schema.Respond(schema.Empty{}, c)
}

// @Summary Authorization
// @Tags users
// @Accept json
// @Produce json
// @Param data body schema.AuthorizationSchema true "AuthorizationSchema"
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /user/auth [post]
func (a *AuthenticationRouter) AuthHandler(ctx *gin.Context) error {
	var payload *schema.AuthorizationSchema

	if err := ctx.BindJSON(&payload); err != nil {
		return err
	}

	err := a.Services.Auth.Authorization(ctx, payload)
	if err != nil {
		return schema.Respond(err, ctx)
	}

	return schema.Respond(schema.Empty{}, ctx)
}

// @Summary Logout
// @Tags users
// @Accept json
// @Produce json
// @Failure 400 {object} object
// @Failure 500 {object} object
// @Router /user/logout [get]
func (a *AuthenticationRouter) LogoutHandler(ctx *gin.Context) error {

	err := a.Services.Auth.Logout(ctx)
	if err != nil {
		return err
	}

	return schema.Respond(schema.Empty{}, ctx)
}
