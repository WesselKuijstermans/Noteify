package repositories

import (
	"github.com/gin-gonic/gin"
	"noteify-api/database"
	"noteify-api/models"
)

func CreateNote(c *gin.Context) {
	var note models.Note
	c.BindJSON(&note)
	if note.Title == "" || note.Content == "" || note.UserId == "" {
		c.JSON(400, gin.H{
			"error": "Title, content and user are required",
		})
		return
	}
	database.DB.Create(&note)
	c.JSON(200, note)
}
