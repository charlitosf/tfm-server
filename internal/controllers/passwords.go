package controllers

import "github.com/gin-gonic/gin"

// Get all passwords handler
func GetPasswords(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Get all passwords from website handler
func GetPasswordsByWebsite(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Create a password on a website handler
func CreatePassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Get a password handler
func GetPassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Update a password handler
func UpdatePassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Delete a password handler
func DeletePassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
