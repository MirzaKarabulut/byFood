package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProcessURL(c *gin.Context) {
	var req struct {
		URL       string `json:"url" binding:"required"`
		Operation string `json:"operation" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Parse the URL
	parsedURL, err := url.Parse(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// Define the operation functions
	canonicalize := func(parsedURL *url.URL) string {
		// Remove query parameters and trailing slashes
		parsedURL.RawQuery = ""
		parsedURL.Fragment = ""
		parsedURL.Path = strings.TrimRight(parsedURL.Path, "/")
		return parsedURL.String()
	}

	redirect := func(parsedURL *url.URL) string {
		// Ensure the domain is www.byfood.com and convert the URL to lowercase
		parsedURL.Host = "www.byfood.com"
		return strings.ToLower(parsedURL.String())
	}

	// Process the URL based on the operation
	var processedURL string
	if req.Operation == "canonical" {
		processedURL = canonicalize(parsedURL)
	} else if req.Operation == "redirection" {
		processedURL = redirect(parsedURL)
	} else if req.Operation == "all" {
		canonical := canonicalize(parsedURL)
		parsedURL, _ = url.Parse(canonical) // re-parse the canonical URL
		processedURL = redirect(parsedURL)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"processed_url": processedURL})
}
