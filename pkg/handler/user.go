package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUser(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	/* Firstname string         `json:"firstname" gorm:"not null"`
	Lastname  string         `json:"lastname" gorm:"not null"`
	Username  string         `json:"username" gorm:"not null,unique"`
	Email     string         `json:"email" gorm:"not null,unique"`
	Password  string         `json:"-" gorm:"not null"` */
}
