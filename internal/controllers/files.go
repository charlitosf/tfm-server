package controllers

import "github.com/gin-gonic/gin"

// Get all files handler
func GetFiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Create a file handler
func CreateFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Get a file handler
func GetFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Update a file handler
func UpdateFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Delete a file handler
func DeleteFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
