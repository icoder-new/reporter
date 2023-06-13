package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icoder-new/reporter/middlewares"
)

type Routes struct{}

func (r *Routes) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
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
			auth.POST("/register")
			auth.POST("/login")
		}

		user := api.Group("/user")
		{
			user.GET("/")
			user.PUT("/")
			user.GET("/restore")
			user.DELETE("/delete")
		}

		account := api.Group("/account")
		{
			account.GET("/")
			account.GET("/:id")
			account.POST("/")
			account.PUT("/:id")
			account.DELETE("/")
		}

		report := api.Group("/report")
		{
			report.GET("/")
			report.GET("/:id")
			report.GET("/:transaction_id")
			report.GET("/:user_id")
		}

		transaction := api.Group("/transaction")
		{
			transaction.GET("/")
			transaction.POST("/")
			transaction.PATCH("/")
		}
	}

	return router
}
