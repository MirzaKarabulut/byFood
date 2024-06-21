package controllers

import (
	"github.com/gin-gonic/gin"
)

// ErrorHandler handles errors and sends appropriate responses
func ErrorHandler(c *gin.Context, statusCode int, err error, message string, detail string) {
    c.JSON(statusCode, gin.H{
        "error":   err.Error(),
        "message": message,
        "detail":  detail,
    })
    c.Abort()
}