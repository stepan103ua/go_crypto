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
	GetPost(postId int) (models.Post, error)
	DeletePost(postId, userId int) error
	UpdatePost(title, description, cryptoCurrency, image_url string, postId, userId int) error
}

type Comments interface {
	GetAllCommentsByPostId(postId int) ([]models.Comment, error)
	CreateComment(comment models.Comment) (int, error)
	DeleteComment(commentId, userId int) error
	UpdateComment(comment string, commentId, userId int) error
	GetCommentsCountByPostId(postId int) (int, error)
}

type Replies interface {
	CreateReply(reply models.Reply) (int, error)
	GetRepliesByCommentId(commentId int) ([]models.Reply, error)
}

type Users interface {
	GetUserById(userId int) (models.UserResponse, error)
	UpdateUser(user models.UserUpdate, userId int) error
}

type Likes interface {
	ToggleLike(postId, userId int) error
	IsLiked(postId, userId int) (bool, error)
	GetLikesCountByPostId(postId int) (int, error)
}

type Followers interface {
	ToggleFollow(followerId, userId int) error
	IsFollowing(followerId, userId int) (bool, error)
	GetFollowersCount(userId int) (int, error)
	GetFollowingCount(userId int) (int, error)
}

type Watchlists interface {
	CreateWatchlist(watchlist models.Watchlist) error
	GetAllUserWatchlists(userId int) ([]models.Watchlist, error)
}

type Service struct {
	Authorization
	Posts
	Comments
	Replies
	Users
	Likes
	Followers
	Watchlists
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Posts:         NewPostsService(repository.Posts),
		Comments:      NewCommentsService(repository.Comments),
		Replies:       NewRepliesService(repository.Replies),
		Users:         NewUsersService(repository.Users),
		Likes:         NewLikesService(repository.Likes),
		Followers:     NewFollowersService(repository.Followers),
		Watchlists:    NewWatchlistsService(repository.Watchlists),
	}
}
