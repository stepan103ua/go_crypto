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
