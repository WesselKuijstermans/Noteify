package routers

import (
	"github.com/gin-gonic/gin"
	"noteify-api/handlers"
	"noteify-api/middleware"
)

func SetupNoteRouter(router *gin.Engine) {
	notes := router.Group("/notes").Use(middleware.AuthMiddleWare())
	{
		notes.GET("", handlers.GetNotes)
		notes.GET("/:id", handlers.GetNote)
		notes.POST("", handlers.CreateNote)
		notes.PUT("/:id", handlers.UpdateNote)
		notes.DELETE("/:id", handlers.DeleteNote)
	}
}
