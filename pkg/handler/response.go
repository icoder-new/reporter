package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/icoder-new/reporter/logger"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResonse(c *gin.Context, status int, message string) {
	logger.Error.Printf(message)
	c.AbortWithStatusJSON(status, errorResponse{
		Message: message,
	})
}
