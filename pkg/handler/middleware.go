package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authoriztionHeader = "Authorization"
	userIdContext      = "userId"
)

func (h *Handler) userIndentity(c *gin.Context) {
	header := c.GetHeader(authoriztionHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid header")
		return
	}

	userId, err := h.service.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIdContext, userId)
}
