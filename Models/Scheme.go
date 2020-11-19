package Models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

type Song struct {
	gorm.Model
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

func (b *Book) TableName() string {
	return "book"
}

func (s *Song) TableName() string {
	return "songs"
}
