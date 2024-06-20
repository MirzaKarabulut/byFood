package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandler handles errors and sends appropriate responses
func ErrorHandler(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
