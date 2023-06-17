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
			user.GET("/restore", h.RestoreUser)
			user.DELETE("/delete", h.DeleteUser)
			user.PATCH("/change", h.ChangePictureUser)
			user.PATCH("/upload", h.UploadUserPicture)
		}

		account := api.Group("/account", h.UserIdentity)
		{
			account.GET("/", h.GetAllAccounts)
			account.GET("/:id", h.GetAccount)
			account.POST("/", h.CreateAccount)
			account.PUT("/:id", h.UpdateAccount)
			account.GET("/:id/restore", h.RestoreAccount)
			account.DELETE("/:id", h.DeleteAccount)
			account.PATCH("/:id/change", h.ChangePictureAccount)
			account.PATCH("/:id/upload", h.UploadAccountPicture)
		}

		transaction := api.Group("/transaction", h.UserIdentity)
		{
			transaction.GET("/")
			transaction.POST("/")
			transaction.PATCH("/")
		}

		report := api.Group("/report", h.UserIdentity)
		{
			report.GET("/")
			// report.GET("/:id")
			// report.GET("/:transaction_id")
			// report.GET("/:user_id")
		}
	}

	return router
}
