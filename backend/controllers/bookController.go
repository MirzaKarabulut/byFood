package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetBooks             godoc
// @Summary      Get books array
// @Description  Responds with the list of all books as JSON.
// @Tags         books
// @Produce      json
// @Success      200  {array}  models.Book
// @Router       /books [get]
func GetAllBooks(c *gin.Context) {
    var book []models.Book
    result := initializers.DB.Find(&book)
    if result.Error != nil {
        ErrorHandler(c, http.StatusInternalServerError, result.Error, "Failed to retrieve books", "Database error")
        return
    }

    LogEvent(c, "Get All Books", "Book", "Success", "Retrieved all books")
    LogRequest(c, "200")

    c.JSON(http.StatusOK, book)
}

// GetBooks             godoc
// @Summary      Get books array
// @Description  Respond with id of book
// @Tags         books
// @Produce      json
// @Param        id  path  string  true  "Book ID"
// @Success      200  {array}  models.Book
// @Router       /books/{id} [get]
func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    result := initializers.DB.First(&book, id)
    if result.Error != nil {
        ErrorHandler(c, http.StatusNotFound, result.Error, "Book not found", "No book exists with the given ID")
        return
    }


    LogEvent(c, "Get Book By ID", "Book", "Success", "Retrieved book by ID")
    LogRequest(c, "200")

    c.JSON(http.StatusOK, book)
}

// CreateBook           godoc
// @Summary      Create a new book
// @Description  Create a new book with the given details.
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.Book  true  "Book JSON"
// @Success      201  {object}  models.Book
// @Router       /books [post] 
func CreateBook(c *gin.Context) {
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


    LogEvent(c, "Create Book", "Book", "Success", "Book created successfully")
    LogRequest(c, "201")

    c.JSON(http.StatusCreated, book)
}


// UpdateBook           godoc
// @Summary      Update a book
// @Description  Update the details of a book with the given ID.
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Book
// @Param        book  body      models.Book  true  "Book JSON"
// @Router       /books/{id} [put]
// @Param        id  path  string  true  "Book ID"
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


    LogEvent(c, "Update Book", "Book", "Success", "Book updated successfully")
    LogRequest(c, "200")

    c.JSON(http.StatusOK, book)
}

// DeleteBook           godoc
// @Summary      Delete a book
// @Description  Delete the book with the given ID.
// @Tags         books
// @Produce      json
// @Param        id  path  string  true  "delete by ID"
// @Success      200  {object}  models.Book
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    result := initializers.DB.Delete(&models.Book{}, id)
    if result.Error != nil {
        ErrorHandler(c, http.StatusInternalServerError, result.Error, "Failed to delete book", "Database error")
        return
    }


    LogEvent(c, "Delete Book", "Book", "Success", "Book deleted successfully")
    LogRequest(c, "200")

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully", "id": id})
}
