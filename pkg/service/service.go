package service

import (
	"crypto_app/pkg/models"
	"crypto_app/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateJWT(email, password string) (string, error)
	ParseToken(jwt string) (int, error)
}

type Posts interface {
	CreatePost(post models.Post) (int, error)
	GetAllPosts() ([]models.Post, error)
	GetAllPostsByUserId(userId int) ([]models.Post, error)
}

type Comments interface {
	GetAllCommentsByPostId(postId int) ([]models.Comment, error)
}

type Service struct {
	Authorization
	Posts
	Comments
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Posts:         NewPostsService(repository.Posts),
		Comments:      NewCommentsService(repository.Comments),
	}
}
