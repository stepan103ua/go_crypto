package handler

import (
	"crypto_app/pkg/models"
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

func (h *Handler) updateUser(c *gin.Context) {
	var user models.UserUpdate

	err := c.BindJSON(&user)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	err = h.service.UpdateUser(user, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}

func (h *Handler) getAuthenticatedUser(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	user, err := h.service.GetUserById(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
