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

func (h *Handler) getFollowersCount(c *gin.Context) {
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

	count, err := h.service.GetFollowersCount(userIdInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong user id")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"count": count,
	})
}

func (h *Handler) getFollowingCount(c *gin.Context) {
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

	count, err := h.service.GetFollowingCount(userIdInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong user id")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"count": count,
	})
}

func (h *Handler) isFollowing(c *gin.Context) {
	followerId := c.Param("userId")

	if followerId == "" {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong user id")
		return
	}

	followerIdInt, err := strconv.Atoi(followerId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong user id")
		return
	}

	userId, err := h.getUserId(c)

	if err != nil {
		return
	}

	isFollowing, err := h.service.IsFollowing(followerIdInt, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong user id")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"isFollowing": isFollowing,
	})
}
