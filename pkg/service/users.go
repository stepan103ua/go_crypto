package service

import (
	"crypto_app/pkg/models"
	"crypto_app/pkg/repository"
)

type UsersService struct {
	repository repository.Users
}

func NewUsersService(repository repository.Users) *UsersService {
	return &UsersService{repository: repository}
}

func (s *UsersService) GetUserById(userId int) (models.UserResponse, error) {
	return s.repository.GetUserById(userId)
}

func (s *UsersService) UpdateUser(user models.UserUpdate, userId int) error {
	if user.Name != "" {
		err := s.repository.UpdateUsername(user.Name, userId)

		if err != nil {
			return err
		}
	}
	if user.Email != "" {
		err := s.repository.UpdateEmail(user.Email, userId)

		if err != nil {
			return err
		}
	}
	if user.Password != "" {
		user.Password = generatePasswordHash(user.Password)
		err := s.repository.UpdatePassword(user.Password, userId)

		if err != nil {
			return err
		}
	}
	if user.About != "" {
		err := s.repository.UpdateAbout(user.About, userId)

		if err != nil {
			return err
		}
	}

	return nil
}
