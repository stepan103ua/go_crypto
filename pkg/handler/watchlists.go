package handler

import (
	"crypto_app/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createWatchlist(c *gin.Context) {
	var watchlist models.Watchlist

	err := c.BindJSON(&watchlist)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	watchlist.UserId = userId

	err = h.service.CreateWatchlist(watchlist)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully created",
	})
}

func (h *Handler) getAllUserWatchlists(c *gin.Context) {
	userId, err := h.getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userWatchlists, err := h.service.GetAllUserWatchlists(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userWatchlists)
}
