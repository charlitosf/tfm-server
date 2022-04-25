package controllers

import (
	"charlitosf/tfm-server/internal/services"
	"charlitosf/tfm-server/pkg/httptypes"

	"github.com/gin-gonic/gin"
)

// Get all passwords handler
func GetPasswords(c *gin.Context) {
	// Get user
	user := c.MustGet("username").(string)
	// Get passwords
	passwords, err := services.GetAllPasswords(user)
	if err != nil {
		c.JSON(400, httptypes.GenericErrorResponse{Error: &httptypes.Error{Message: err.Error()}})
	} else {
		c.JSON(200, passwords)
	}
}

// Get all passwords from website handler
func GetPasswordsByWebsite(c *gin.Context) {
	// Get website
	website := c.Param("website")
	// Get user
	user := c.MustGet("username").(string)
	// Get passwords
	passwords, err := services.GetPasswords(user, website)
	if err != nil {
		c.JSON(400, httptypes.GenericErrorResponse{Error: &httptypes.Error{Message: err.Error()}})
	} else {
		c.JSON(200, passwords)
	}
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
	// Get user
	user := c.MustGet("username").(string)
	// Get website
	website := c.Param("website")
	// Get username
	username := c.Param("username")
	// Get password
	password, err := services.GetPassword(user, website, username)
	if err != nil {
		c.JSON(400, httptypes.GetPasswordResponse{Error: &httptypes.Error{Message: err.Error()}})
	} else {
		c.JSON(200, httptypes.GetPasswordResponse{Password: password})
	}
}

// Update a password handler
func UpdatePassword(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Delete a password handler
func DeletePassword(c *gin.Context) {
	// Get user
	user := c.MustGet("username").(string)
	// Get website
	website := c.Param("website")
	// Get username
	username := c.Param("username")
	// Delete password
	err := services.DeletePassword(user, website, username)
	if err != nil {
		c.JSON(400, httptypes.GenericErrorResponse{Error: &httptypes.Error{Message: err.Error()}})
	} else {
		c.JSON(200, httptypes.GenericEmptyResponse{})
	}
}
