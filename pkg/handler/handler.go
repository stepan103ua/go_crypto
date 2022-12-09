package handler

import (
	"crypto_app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signUp", h.signUp)
		auth.POST("/signIn", h.signIn)
	}
	api := router.Group("/api", h.userIndentity)
	{
		posts := api.Group("/posts")
		{
			posts.GET("/", h.getAllPosts)
			posts.GET("/:postId", h.getPost)
			posts.DELETE("/:postId", h.deletePost)
			posts.POST("/:postId/toggleLike", h.toggleLike)
			posts.GET("/:postId/likes", h.getLikesCountByPostId)
			posts.PUT("/:postId", h.updatePost)
			posts.GET("/myAll", h.getAllUserPosts)
			posts.POST("/new", h.createPost)
			posts.GET("/users/:userId", h.getAllPostsByUserId)
		}
		comments := api.Group("/comments")
		{
			comments.GET("/posts/:postId", h.getAllCommentsByPostId)
			comments.GET("/posts/:postId/count", h.getCommentsCountByPostId)
			comments.DELETE("/:commentId", h.deleteComment)
			comments.PUT("/:commentId", h.updateComment)
			comments.POST("/new", h.createComment)

			comments.POST("/:commentId/replies/new", h.createReply)
			comments.GET("/:commentId/replies", h.getRepliesByCommentId)
		}
		users := api.Group("/users")
		{
			users.GET("/:userId", h.getUserById)
			users.PUT("/update", h.updateUser)
			users.GET("/me", h.getAuthenticatedUser)
			users.POST("/:followerId/toggleFollow", h.toggleFollow)
			users.GET("/:userId/followersCount", h.getFollowersCount)
			users.GET("/:userId/followingCount", h.getFollowingCount)
			users.GET("/:userId/isFollowing", h.isFollowing)
		}
		watchlists := api.Group("/watchlists")
		{
			watchlists.POST("/new", h.createWatchlist)
		}
	}

	return router
}
