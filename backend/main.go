package main

import (
	"backend/controllers"
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
    r.GET("/books", controllers.GetAllBooks)
    r.POST("/books", controllers.CreateBook)
    r.GET("/books/:id", controllers.GetBookByID)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.DELETE("/books/:id", controllers.DeleteBook)
    r.POST("/process-url",controllers.ProcessURL)
    r.Run()
}