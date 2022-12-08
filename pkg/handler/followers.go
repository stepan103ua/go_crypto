package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) toggleFollow(c *gin.Context) {
	followerId := c.Param("followerId")

	if followerId == "" {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong follower id")
		return
	}

	followerIdInt, err := strconv.Atoi(followerId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong follower id")
		return
	}

	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	err = h.service.ToggleFollow(followerIdInt, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	isFollowing, err := h.service.IsFollowing(followerIdInt, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isFollowing": isFollowing,
	})
}
