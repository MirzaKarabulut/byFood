package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetAllBooks(c *gin.Context) {
    // Retrieve all books from the database
    var books []models.Book
    result := initializers.DB.Find(&books)
    if result.Error != nil {
        ErrorHandler(c, http.StatusInternalServerError, result.Error, "Failed to retrieve books", "Database error")
        return
    }

    // Log event
    LogEvent(c, "Get All Books", "Book", "Success", "Retrieved all books")
    // Log the request
    LogRequest(c, "200")

    // Return the books as JSON response
    c.JSON(http.StatusOK, books)
}


func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    result := initializers.DB.First(&book, id)
    if result.Error != nil {
        ErrorHandler(c, http.StatusNotFound, result.Error, "Book not found", "No book exists with the given ID")
        return
    }

    // Log event
    LogEvent(c, "Get Book By ID", "Book", "Success", "Retrieved book by ID")
    // Log the request
    LogRequest(c, "200")

    c.JSON(http.StatusOK, book)
}


func CreateBook(c *gin.Context) {
    // Parse the request body
    var body struct {
        Title       string `json:"title" validate:"required"`
        Author      string `json:"author" validate:"required"`
        Description string `json:"description" validate:"required"`
        ReleaseDate string `json:"releaseDate" validate:"required"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        ErrorHandler(c, http.StatusBadRequest, err, "Invalid JSON format", "Check the request body for errors")
        return
    }
    validate := validator.New()
		err := validate.Struct(body);
    if  err != nil {
        ErrorHandler(c, http.StatusBadRequest, err, "Validation error", "Check the input fields")
        return
    }

    book := models.Book{
        Title:       &body.Title,
        Author:      &body.Author,
        Description: &body.Description,
        ReleaseDate: &body.ReleaseDate,
    }

    tx := initializers.DB.Begin()
    if err := tx.Create(&book).Error; err != nil {
        tx.Rollback()
        ErrorHandler(c, http.StatusInternalServerError, err, "Failed to create book", "Database error")
        return
    }
    tx.Commit()

    // Log event
    LogEvent(c, "Create Book", "Book", "Success", "Book created successfully")
    // Log the request
    LogRequest(c, "201")

    c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var body struct {
        Title       string `json:"title"`
        Author      string `json:"author"`
        Description string `json:"description"`
        ReleaseDate string `json:"releaseDate"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        ErrorHandler(c, http.StatusBadRequest, err, "Invalid JSON format", "Check the request body for errors")
        return
    }

    var book models.Book
    result := initializers.DB.First(&book, id)
    if result.Error != nil {
        ErrorHandler(c, http.StatusNotFound, result.Error, "Book not found", "No book exists with the given ID")
        return
    }

    result = initializers.DB.Model(&book).Updates(models.Book{
        Title:       &body.Title,
        Author:      &body.Author,
        Description: &body.Description,
        ReleaseDate: &body.ReleaseDate,
    })
    if result.Error != nil {
        ErrorHandler(c, http.StatusInternalServerError, result.Error, "Failed to update book", "Database error")
        return
    }

    // Log event
    LogEvent(c, "Update Book", "Book", "Success", "Book updated successfully")
    // Log the request
    LogRequest(c, "200")

    c.JSON(http.StatusOK, book)
}


func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    result := initializers.DB.Delete(&models.Book{}, id)
    if result.Error != nil {
        ErrorHandler(c, http.StatusInternalServerError, result.Error, "Failed to delete book", "Database error")
        return
    }

    // Log event
    LogEvent(c, "Delete Book", "Book", "Success", "Book deleted successfully")
    // Log the request
    LogRequest(c, "200")

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully", "id": id})
}

