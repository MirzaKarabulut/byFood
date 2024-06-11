package controllers

import (
	"backend/initializers"
	"backend/models"

	"github.com/gin-gonic/gin"
)


func GetAllBooks(c *gin.Context)  {
	var books []models.Book
 initializers.DB.Find(&books)
	c.JSON(200, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	initializers.DB.Create(&book)
	c.JSON(200, book)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	initializers.DB.First(&book, id)
	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	initializers.DB.First(&book, id)
	c.BindJSON(&book)
	initializers.DB.Save(&book)
	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	initializers.DB.First(&book, id)
	initializers.DB.Delete(&book)
	c.JSON(200, book)
}