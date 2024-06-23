package main

import (
	"backend/controllers"
	_ "backend/docs"
	"backend/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
    initializers.LoadEnvVars()
    initializers.Connect()
}

// @title ByFood API
// @version 1.0
// @description This is a simple API for a book store
// @summary This is a simple API for a book store
// @host localhost:8080
// @BasePath /


func main() {
    r := gin.Default()
    r.Use(cors.Default())
    r.GET("/books", controllers.GetAllBooks)
    r.POST("/books", controllers.CreateBook)
    r.GET("/books/:id", controllers.GetBookByID)
    r.PUT("/books/:id", controllers.UpdateBook)
    r.DELETE("/books/:id", controllers.DeleteBook)
    r.POST("/process-url",controllers.ProcessURL)
    
    r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run()
}