package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/icoder-new/reporter/pkg/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(h.CORSMiddleware())
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping - pong",
		})
	})

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.SignUp)
			auth.POST("/login", h.signIn)
		}

		user := api.Group("/user", h.UserIdentity)
		{
			user.GET("/", h.GetUser)
			user.PUT("/", h.UpdateUser)
			user.GET("/restore")
			user.DELETE("/delete")
		}

		account := api.Group("/account", h.UserIdentity)
		{
			account.GET("/")
			account.GET("/:id")
			account.POST("/")
			account.PUT("/:id")
			account.DELETE("/")
		}

		report := api.Group("/report", h.UserIdentity)
		{
			report.GET("/")
			// report.GET("/:id")
			// report.GET("/:transaction_id")
			// report.GET("/:user_id")
		}

		transaction := api.Group("/transaction", h.UserIdentity)
		{
			transaction.GET("/")
			transaction.POST("/")
			transaction.PATCH("/")
		}
	}

	return router
}
