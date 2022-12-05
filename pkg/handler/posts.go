package handler

import (
	"crypto_app/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetAllPostsResponse struct {
	Data []models.Post `json:"posts"`
}

func (h *Handler) createPost(c *gin.Context) {
	var input models.Post

	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	input.UserId = userId

	id, createPostErr := h.service.CreatePost(input)

	if createPostErr != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) getAllUserPosts(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}
	posts, err := h.service.GetAllPostsByUserId(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	jsonPosts(c, posts)
}

func (h *Handler) getAllPosts(c *gin.Context) {

	posts, err := h.service.GetAllPosts()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	jsonPosts(c, posts)
}

func (h *Handler) getAllPostsByUserId(c *gin.Context) {
	userId := c.Param("userId")

	if userId == "" {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong user id")
		return
	}

	userIdInt, err := strconv.Atoi(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong user id")
	}

	posts, err := h.service.GetAllPostsByUserId(userIdInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	jsonPosts(c, posts)
}

func jsonPosts(c *gin.Context, posts []models.Post) {
	if posts == nil {
		c.JSON(http.StatusOK,
			GetAllPostsResponse{Data: make([]models.Post, 0)},
		)
		return
	}

	c.JSON(http.StatusOK,
		GetAllPostsResponse{Data: posts},
	)
}
