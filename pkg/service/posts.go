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

func (s *PostsService) DeletePost(postId, userId int) error {
	return s.repository.DeletePost(postId, userId)
}

func (s *PostsService) UpdatePost(title, description, cryptoCurrency, image_url string, postId, userId int) error {
	return s.repository.UpdatePost(title, description, cryptoCurrency, image_url, postId, userId)
}

func (s *PostsService) GetPost(postId int) (models.Post, error) {
	return s.repository.GetPost(postId)
}
