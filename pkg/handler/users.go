package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserById(c *gin.Context) {
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

	user, err := h.service.GetUserById(userIdInt)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
