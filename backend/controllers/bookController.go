package controllers

import (
	"backend/initializers"
	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAllBooks(c *gin.Context) {
	// Log event
	LogEvent(c, "Get All Books", "Book", "Success", "Retrieved all books")
	// Log the request
	LogRequest(c, "200")

	// Retrieve all books from the database
	var books []models.Book
	result := initializers.DB.Find(&books)
	if result.Error != nil {
		ErrorHandler(c, result.Error)
		return
	}

	// Return the books as JSON response
	c.JSON(200, books)
}

func GetBookByID(c *gin.Context) {
	// Log event
	LogEvent(c, "Get Book By ID", "Book", "Success", "Retrieved book by ID")

	// Log the request
	LogRequest(c, "200")

	id := c.Param("id")
	var book models.Book
	result := initializers.DB.First(&book, id)
	if result.Error != nil {
		ErrorHandler(c, result.Error)
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	// Parse the request body
	var body struct {
		Title       string `json:"Title" validate:"required"`
		Author      string `json:"Author" validate:"required"`
		Description string `json:"Description" validate:"required"`
		ReleaseDate string `json:"ReleaseDate" validate:"required,datetime=2006-01-02"`
	}
	err := c.Bind(&body)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		LogEvent(c, "Create Book", "Book", "Validation", err.Error())
		LogRequest(c, "400")
		ErrorHandler(c, err)
		return
	}

	tx := initializers.DB.Begin()

	book := models.Book{Title: &body.Title, Author: &body.Author, Description: &body.Description, ReleaseDate: &body.ReleaseDate}
	result := tx.Create(&book)
	if result.Error != nil {
		LogEvent(c, "Create Book", "Book", "Database", result.Error.Error())
		LogRequest(c, "400")
		tx.Rollback()
		ErrorHandler(c, result.Error)
		return
	}

	tx.Commit()
	LogEvent(c, "Create Book", "Book", "Success", "Book created successfully")
	LogRequest(c, "200")

	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title       string
		Author      string
		Description string
		ReleaseDate string
	}
	err := c.Bind(&body)
	if err != nil {
		ErrorHandler(c, err)
		return
	}

	var book models.Book
	result := initializers.DB.First(&book, id)
	if result.Error != nil {
		ErrorHandler(c, result.Error)
		return
	}

	result = initializers.DB.Model(&book).Updates(models.Book{Title: &body.Title, Author: &body.Author, Description: &body.Description, ReleaseDate: &body.ReleaseDate})
	if result.Error != nil {
		ErrorHandler(c, result.Error)
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	// Log event
	LogEvent(c, "Delete Book", "Book", "Success", "Book deleted successfully")

	// Log the request
	LogRequest(c, "200")

	id := c.Param("id")
	result := initializers.DB.Delete(&models.Book{}, id)
	if result.Error != nil {
		ErrorHandler(c, result.Error)
		return
	}

	c.JSON(200, gin.H{"message": "Book deleted successfully", "id": id})
}
