package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"regexp"
)

type User struct {
	gorm.Model
	ID    string `json:"ID" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique <-:create"`
	Name  string `json:"name"`
	Notes []Note `json:"notes" gorm:"foreignKey:user_id"`
}

// HasValidEmail checks if the email is valid
func (u *User) HasValidEmail() bool {
	// Regular expression pattern for email validation
	// This pattern allows for most common email formats
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression pattern
	regex := regexp.MustCompile(emailPattern)

	// Check if the email matches the pattern
	return regex.MatchString(u.Email)
}

// GenerateUUID generates a random uuid for the user
func (u *User) GenerateUUID() {
	randomID, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	u.ID = randomID.String()
}
