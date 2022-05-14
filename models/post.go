package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string `json:"title"`
	AuthorID int    `json:"author_id"`
	Author   Author `json:"author"`
	Text     string `json:"text"`
}

type CreatePost struct {
	Title  string `json:"title" binding:"required"`
	Author Author `json:"author" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

type UpdatePost struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
