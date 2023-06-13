package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/icoder-new/reporter/models"
)

type signInData struct {
	Username string `json:"username" binding:"require"`
	Password string `json:"password" binding:"require"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var request models.User

	if err := c.BindJSON(&request); err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(request)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var request signInData

	if err := c.BindJSON(&request); err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.GenerateToken(request.Username, request.Password)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
