package services

import (
	"final/models"
	"final/persistence"
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
		return user, err
	}

	return user, nil
}
