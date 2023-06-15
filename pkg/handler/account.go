package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type accountData struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (h *Handler) GetAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	account, err := h.service.GetAccount(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"account": account,
	})
}

func (h *Handler) GetAllAccounts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accounts, err := h.service.GetAllAccounts(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"accounts": accounts,
	})
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var request accountData
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateAccount(userId, request.Name, request.Balance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "success",
	})
}

func (h *Handler) UpdateAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	var request accountData
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&request); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	account, err := h.service.UpdateAccount(id, userId, request.Name, request.Balance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    account,
	})
}

func (h *Handler) DeleteAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = h.service.DeleteAccount(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h *Handler) RestoreAccount(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = h.service.RestoreAccount(id, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
