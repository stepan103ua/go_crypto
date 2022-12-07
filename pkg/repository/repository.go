package repository

import (
	"crypto_app/pkg/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(email, password string) (models.User, error)
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

type Users interface {
	GetUserById(userId int) (models.UserResponse, error)
}

type Repository struct {
	Authorization
	Posts
	Comments
	Replies
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Posts:         NewPostsPostgres(db),
		Comments:      NewCommentsPostgres(db),
		Replies:       NewRepliesPostgres(db),
		Users:         NewUsersPostgres(db),
	}
}
