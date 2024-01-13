package schema

type AuthorizationSchema struct {
	LoginOrEmail string `json:"login_or_email" validate:"required"`
	Password     string `json:"password" validate:"required"`
}
