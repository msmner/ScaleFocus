package services

import (
	"final/models"
	"final/persistence"
	"log"
)

type UserService struct {
	userRepository *persistence.UserRepository
}

func NewUserService(ur *persistence.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (us *UserService) GetUser(username string) (models.User, error) {
	log.Printf("username in getuser service is %s", username)
	user, err := us.userRepository.GetUser(username)
	if err != nil {
		log.Printf("error in getuser service is %v", err)
		return user, err
	}

	return user, nil
}
