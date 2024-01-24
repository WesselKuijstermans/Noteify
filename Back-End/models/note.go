package models

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserId  string `json:"user_id"`
}
