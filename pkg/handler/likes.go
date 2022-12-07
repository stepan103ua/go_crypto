package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) toggleLike(c *gin.Context) {
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

	err = h.service.ToggleLike(postIdInt, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})

}
