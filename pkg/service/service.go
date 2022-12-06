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
	DeletePost(postId, userId int) error
	UpdatePost(title, description, cryptoCurrency, image_url string, postId, userId int) error
}

type Comments interface {
	GetAllCommentsByPostId(postId int) ([]models.Comment, error)
	CreateComment(comment models.Comment) (int, error)
	DeleteComment(commentId, userId int) error
	UpdateComment(comment string, commentId, userId int) error
}

type Replies interface {
	CreateReply(reply models.Reply) (int, error)
	GetRepliesByCommentId(commentId int) ([]models.Reply, error)
}

type Service struct {
	Authorization
	Posts
	Comments
	Replies
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Posts:         NewPostsService(repository.Posts),
		Comments:      NewCommentsService(repository.Comments),
		Replies:       NewRepliesService(repository.Replies),
	}
}
