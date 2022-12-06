package handler

import (
	"crypto_app/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCommentsResponse struct {
	Data []models.Comment `json:"comments"`
}

func (h *Handler) getAllCommentsByPostId(c *gin.Context) {
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

	comments, err := h.service.GetAllCommentsByPostId(postIdInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetCommentsResponse{Data: comments})
}

func (h *Handler) createComment(c *gin.Context) {
	var input models.Comment

	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	input.OwnerId = userId

	commentId, err := h.service.CreateComment(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": commentId,
	})
}

func (h *Handler) deleteComment(c *gin.Context) {
	commentId := c.Param("commentId")

	if commentId == "" {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong comment id")
		return
	}

	commentIdInt, err := strconv.Atoi(commentId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong comment id")
		return
	}

	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	err = h.service.DeleteComment(commentIdInt, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted",
	})

}
