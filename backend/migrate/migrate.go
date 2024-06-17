package main

import (
	"backend/initializers"
	"backend/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.Connect()
}

func main() {
	 err := initializers.DB.AutoMigrate(&models.Book{})
	 if err != nil {
		return
	 }
}