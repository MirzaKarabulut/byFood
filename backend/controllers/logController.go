package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// LogRequest logs details of a request to the database
func LogRequest(c *gin.Context, responseStatus string) {
	logBody := models.RequestLog{
		UserAgent: c.Request.UserAgent(),
		SessionID: c.Request.Header.Get("SessionID"),
		DateTime:  time.Now().Format("2006-01-02 15:04:05"),
		Request:   c.Request.Method + " " + c.Request.URL.Path,
		Response:  responseStatus,
	}
	tx := initializers.DB.Begin()
	result := tx.Create(&logBody)
	if result.Error != nil {
		tx.Rollback()
		ErrorHandler(c, http.StatusInternalServerError, result.Error, "Failed to log request", "Database error")
		return
	}
	tx.Commit()
}

// LogEvent logs an event to the database
func LogEvent(c *gin.Context, eventName string, source string, tags string, description string) {
	logBody := models.EventLog{
		EventName:   eventName,
		Source:      source,
		Tags:        tags,
		Description: description,
		UserAgent:   c.Request.UserAgent(),
		SessionID:   c.Request.Header.Get("SessionID"),
		DateTime:    time.Now().Format("2006-01-02 15:04:05"),
	}
	tx := initializers.DB.Begin()
	result := tx.Create(&logBody)
	if result.Error != nil {
		tx.Rollback()
		ErrorHandler(c, http.StatusInternalServerError, result.Error, "Failed to log event", "Database error")
		return
	}
	tx.Commit()
}