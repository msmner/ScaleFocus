package services

import (
	"final/models"
	"final/persistence"
	"fmt"
)

type UserService struct {
	userRepository *persistence.UserRepository
}

func NewUserService(ur *persistence.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (us *UserService) GetUser(username string) (models.User, error) {
	user, err := us.userRepository.GetUser(username)
	if err != nil {
		return user, fmt.Errorf("error getting user: %w", err)
	}

	return user, nil
}
