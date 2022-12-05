package service

import (
	"crypto_app/pkg/models"
	"crypto_app/pkg/repository"
)

type PostsService struct {
	repository repository.Posts
}

func NewPostsService(repository repository.Posts) *PostsService {
	return &PostsService{repository: repository}
}

func (s *PostsService) CreatePost(post models.Post) (int, error) {
	return s.repository.CreatePost(post)
}

func (s *PostsService) GetAllPostsByUserId(userId int) ([]models.Post, error) {
	return s.repository.GetAllPostsByUserId(userId)
}

func (s *PostsService) GetAllPosts() ([]models.Post, error) {
	return s.repository.GetAllPosts()
}
