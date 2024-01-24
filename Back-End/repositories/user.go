package repositories

import (
	"github.com/gin-gonic/gin"
	"log"
	"noteify-api/database"
	"noteify-api/models"
)

func GetUser(c *gin.Context) {
	db := database.DB
	log.Print(db)
	var user models.User
	id := c.Param("id")
	err := db.Preload("Notes").First(&user, id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found!"})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func GetUsers(c *gin.Context) {
	db := database.DB
	var users []models.User
	db.Find(&users)

	c.JSON(200, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	db := database.DB
	var user models.User
	c.BindJSON(&user)

	if user.HasValidEmail() {
		user.GenerateUUID()
		err := db.Create(&user).Error
		if err != nil {
			// duplicate key value violates unique constraint "users_email_key"
			c.JSON(400, gin.H{"error": "Email already exists!"})
		}
	} else {
		c.JSON(400, gin.H{"error": "Invalid email!"})
	}

	c.JSON(201, gin.H{"data": user})
}

func CreateUserFromParams(email string, name string) {
	db := database.DB
	var user models.User
	user.Email = email
	user.Name = name

	if user.HasValidEmail() {
		user.GenerateUUID()
		err := db.Create(&user).Error
		if err != nil {
			// duplicate key value violates unique constraint "users_email_key"
			log.Print("Error: ", err)
		}
	} else {
		log.Print("Invalid email!")
	}
}

func UpdateUser(c *gin.Context) {
	db := database.DB
	var user models.User
	id := c.Param("id")
	err := db.First(&user, id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found!"})
		return
	}
	var updatedUser models.User
	c.BindJSON(&updatedUser)
	db.Model(&user).Updates(updatedUser)
}

func DeleteUser(c *gin.Context) {
	db := database.DB
	var user models.User
	id := c.Param("id")
	err := db.First(&user, id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found!"})
		return
	}
	db.Delete(&user)
}
