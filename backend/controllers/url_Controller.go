package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @Summary Create a new URL
// @Description Create a new URL with the given details.
// @Tags urls
// @Accept  json
// @Produce  json
// @Param url body string true "URL"
// @Success 201 {object} string
// @Router /process-url [post]
func ProcessURL(c *gin.Context) {
	var req struct {
		URL       string `json:"url" binding:"required"`
		Operation string `json:"operation" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	parsedURL, err := url.Parse(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	canonicalize := func(parsedURL *url.URL) string {
		parsedURL.RawQuery = ""
		parsedURL.Fragment = ""
		parsedURL.Path = strings.TrimRight(parsedURL.Path, "/")
		return parsedURL.String()
	}

	redirect := func(parsedURL *url.URL) string {
		parsedURL.Host = "www.byfood.com"
		return strings.ToLower(parsedURL.String())
	}

	var processedURL string
	if req.Operation == "canonical" {
		processedURL = canonicalize(parsedURL)
	} else if req.Operation == "redirection" {
		processedURL = redirect(parsedURL)
	} else if req.Operation == "all" {
		canonical := canonicalize(parsedURL)
		parsedURL, _ = url.Parse(canonical)
		processedURL = redirect(parsedURL)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"processed_url": processedURL})
}
