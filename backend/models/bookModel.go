package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID				 		uint               `json:"id" gorm:"primaryKey"`
	Title       	*string             `json:"title"`
	Author     	 	*string             `json:"author"`
	Description 	*string             `json:"description"`
}


func MigrateBooks(db *gorm.DB) {
	 db.AutoMigrate(&Book{})
}