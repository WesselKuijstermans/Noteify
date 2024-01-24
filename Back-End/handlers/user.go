package handlers

import (
	"github.com/gin-gonic/gin"
	"noteify-api/repositories"
)

func GetUser(c *gin.Context) {
	repositories.GetUser(c)
}

func GetUsers(c *gin.Context) {
	repositories.GetUsers(c)
}

func CreateUser(c *gin.Context) {
	repositories.CreateUser(c)
}

func UpdateUser(c *gin.Context) {
	repositories.UpdateUser(c)
}

func DeleteUser(c *gin.Context) {
	repositories.DeleteUser(c)
}
