package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/icoder-new/reporter/logger"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, status int, message string) {
	logger.Error.Println(message)
	c.AbortWithStatusJSON(status, errorResponse{
		Message: message,
	})
}
