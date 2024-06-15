package controllers

import (
	"backend/initializers"
	"backend/models"
	"fmt"

	"github.com/gin-gonic/gin"
)


func GetAllBooks(c *gin.Context)  {
	var books []models.Book
 initializers.DB.Find(&books)
	c.JSON(200, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	initializers.DB.First(&book, id)
	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	var body struct {
			Title    string
			Author string
			Description string
			ReleaseDate string
	}
	c.Bind(&body)
	fmt.Println(body)

	book := models.Book{Title: &body.Title,Author: &body.Author, Description: &body.Description, ReleaseDate: &body.ReleaseDate}
	initializers.DB.Create(&book)

	c.JSON(200, book)
}


func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var body struct {
			Title    string
			Author string
			Description string
			ReleaseDate string
	}
	c.Bind(&body)
	var book models.Book
	initializers.DB.First(&book, id)
	initializers.DB.Model(&book).Updates(models.Book{Title: &body.Title, Author: &body.Author, Description: &body.Description, ReleaseDate: &body.ReleaseDate})
	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Book{}, id)
	c.JSON(200, id)
}