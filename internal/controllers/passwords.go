package controllers

import (
	"charlitosf/tfm-server/internal/services"
	"charlitosf/tfm-server/pkg/httptypes"

	"github.com/gin-gonic/gin"
)

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
	// Bind request body
	var request httptypes.CreatePasswordRequest
	err := c.BindJSON(&request)
	if err == nil {
		user := c.MustGet("username").(string)
		website := c.Param("website")
		// Create password
		err = services.CreatePassword(user, website, request.Username, request.Password)
		if err != nil {
			c.JSON(400, httptypes.CreatePasswordResponse{Error: &httptypes.Error{Message: err.Error()}})
		} else {
			c.JSON(200, httptypes.CreatePasswordResponse{})
		}
	} else {
		c.JSON(400, httptypes.CreatePasswordResponse{Error: &httptypes.Error{Message: err.Error()}})
	}
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
