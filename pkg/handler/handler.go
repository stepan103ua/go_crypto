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
			posts.GET("/myAll", h.getAllUserPosts)
			posts.POST("/new", h.createPost)
			posts.GET("/users/:userId", h.getAllPostsByUserId)
		}
	}

	return router
}
