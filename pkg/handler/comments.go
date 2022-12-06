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
