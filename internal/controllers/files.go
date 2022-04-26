package controllers

import (
	"charlitosf/tfm-server/internal/dataaccess"
	"charlitosf/tfm-server/pkg/httptypes"

	"github.com/gin-gonic/gin"
)

// Get all files handler
func GetFiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Create a file handler
func CreateFile(c *gin.Context) {
	// Get user from the context
	user := c.MustGet("username").(string)
	// Bind request body to CreateFile struct
	var createFile httptypes.CreateFile
	err := c.BindJSON(&createFile)
	if err == nil {
		// Call dataaccess CreateFile method
		err = dataaccess.CreateFile(user, createFile.Name, createFile.Content)
		if err == nil {
			c.JSON(201, httptypes.GenericResponse{})
		} else {
			c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
		}
	} else {
		c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
	}
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
