package auth

import (
	"final/interfaces"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

type AuthMiddleware struct {
	userService interfaces.IUserService
}

func NewAuthMiddleware(us interfaces.IUserService) *AuthMiddleware {
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
