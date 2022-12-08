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

type UpdatePostRequest struct {
	Title          string `json:"title" binding:"required"`
	Description    string `json:"description" binding:"required"`
	CryptoCurrency string `json:"crypto_currency" binding:"required"`
	ImageUrl       string `json:"image_url" binding:"required"`
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
		return
	}

	posts, err := h.service.GetAllPostsByUserId(userIdInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
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

func (h *Handler) deletePost(c *gin.Context) {
	postId := c.Param("postId")

	if postId == "" {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong post id")
		return
	}

	postIdInt, err := strconv.Atoi(postId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong post id")
		return
	}

	userId, err := h.getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.service.DeletePost(postIdInt, userId)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted",
	})
}

func (h *Handler) updatePost(c *gin.Context) {
	var input UpdatePostRequest

	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	postId := c.Param("postId")

	if postId == "" {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong post id")
		return
	}

	postIdInt, err := strconv.Atoi(postId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong post id")
		return
	}

	userId, err := h.getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.UpdatePost(input.Title, input.Description, input.CryptoCurrency, input.ImageUrl, postIdInt, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully updated",
	})
}

func (h *Handler) getPost(c *gin.Context) {
	postId := c.Param("postId")

	if postId == "" {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong post id")
		return
	}

	postIdInt, err := strconv.Atoi(postId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong post id")
		return
	}

	post, err := h.service.GetPost(postIdInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}
