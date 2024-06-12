package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID				 		uint               `json:"id" gorm:"primaryKey"`
	Title       	*string             `json:"title"`
	Author     	 	*string             `json:"author"`
	ReleaseDate 	*string             `json:"release_date"`
	Description 	*string             `json:"description"`
}
