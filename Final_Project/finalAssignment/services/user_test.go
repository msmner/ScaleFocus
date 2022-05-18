package services

import (
	"final/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	mockUserRepository := MockUserRepository{}
	userService := NewUserService(&mockUserRepository)

	_, err := userService.GetUser("error")
	assert.EqualError(t, err, "error getting user: error")

	expectedUser := models.User{Username: "username", PasswordHash: "hash", ListIds: "16,17"}
	user, _ := userService.GetUser("username")
	assert.Equal(t, expectedUser, user)
}
