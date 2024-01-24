package routers

import (
	"github.com/gin-gonic/gin"
	"noteify-api/handlers"
	"noteify-api/middleware"
)

func SetupUserRouter(router *gin.Engine) {
	users := router.Group("/users").Use(middleware.AuthMiddleWare())
	{
		users.GET("", handlers.GetUsers)
		users.GET("/:id", handlers.GetUser)
		users.POST("", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
}
