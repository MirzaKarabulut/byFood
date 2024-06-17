package controllers

import (
	"backend/initializers"
	"backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func LogRequest(c *gin.Context, responseStatus string) {
	logBody := models.RequestLog{
		UserAgent: c.Request.UserAgent(),
		SessionID: c.Request.Header.Get("SessionID"),
		DateTime:  time.Now().Format("2006-01-02 15:04:05"),
		Request:   c.Request.Method + "" + c.Request.URL.Path,
		Response: responseStatus,
	}
	tx := initializers.DB.Begin()
	result := tx.Create(&logBody)
	if result.Error != nil {
		tx.Rollback()
		c.JSON(400, result.Error)
		return
	}
	tx.Commit()
}

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
		c.JSON(400, result.Error)
		return
	}
	tx.Commit()
}