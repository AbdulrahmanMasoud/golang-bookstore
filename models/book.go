package models

import (
	// "github.com/jinzhu/gorm"
)
type Book struct {
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Author string `json:"author"`
}

type CreateBookInput struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

type UpdateBookInput struct{
	Title string `json:"title"`
	Author string `json:"author"`
}