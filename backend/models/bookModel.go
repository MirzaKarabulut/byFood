package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       	*string             `json:"title" validate:"required"` 
	Author     	 	*string             `json:"author" validate:"required"`
	ReleaseDate 	*string             `json:"release_date" validate:"required" validate:"datetime=2006-01-02"`
	Description 	*string             `json:"description" validate:"required"`
}