package handler

import (
	"crypto_app/pkg/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, createUserErr := h.service.CreateUser(input)

	if createUserErr != nil {
		newErrorResponse(c, http.StatusInternalServerError, createUserErr.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput
	err := c.BindJSON(&input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, tokenError := h.service.GenerateJWT(input.Email, input.Password)

	if tokenError != nil {
		newErrorResponse(c, http.StatusInternalServerError, tokenError.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"jwt": token,
	})
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userIdContext)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "User is not found")
		return 0, errors.New("user is not found")
	}
	idInt, ok := id.(int)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Wrong user id")
		return 0, errors.New("wrong user id")
	}

	return idInt, nil
}
