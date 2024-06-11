package main

import (
	"backend/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
  initializers.LoadEnvVars()
    initializers.Connect()
}

func main() {
    r := gin.Default()
    r.Use(cors.Default())
    r.GET("/books", initializers.GetAllBooks)
    r.POST("/books", initializers.CreateBook)
    r.GET("/books/:id", initializers.GetBookByID)
    r.PUT("/books/:id", initializers.UpdateBook)
    r.DELETE("/books/:id", initializers.DeleteBook)
    r.Run(":8080")
}