package auth

import (
	"final/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

type AuthMiddleware struct {
	userService *services.UserService
}

func NewAuthMiddleware(us *services.UserService) *AuthMiddleware {
	return &AuthMiddleware{userService: us}
}

func (auth *AuthMiddleware) Authenticate(e *echo.Echo) {
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		user, err := auth.userService.GetUser(username)
		if err != nil {
			return false, err
		}

		c.Set("user", username)
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			return false, err
		}

		return true, nil
	}))
}
