package controllers

import "github.com/gin-gonic/gin"

// Update user ('s password) handler
func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Delete user handler
func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
