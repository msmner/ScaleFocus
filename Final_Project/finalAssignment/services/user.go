package services

import (
	"final/interfaces"
	"final/models"
	"fmt"
)

type UserService struct {
	userRepository interfaces.IUserRepository
}

func NewUserService(ur interfaces.IUserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (us *UserService) GetUser(username string) (models.User, error) {
	user, err := us.userRepository.GetUser(username)
	if err != nil {
		return user, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}
