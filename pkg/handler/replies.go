package handler

import (
	"crypto_app/pkg/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetReplies struct {
	Data []models.Reply `json:"replies"`
}

func (h *Handler) createReply(c *gin.Context) {
	var input models.Reply

	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	input.CommentId = commentIdInt
	input.OwnerId = userId

	replyId, err := h.service.CreateReply(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": replyId,
	})
}

func (h *Handler) getRepliesByCommentId(c *gin.Context) {
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

	replies, err := h.service.GetRepliesByCommentId(commentIdInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetReplies{replies})
}
