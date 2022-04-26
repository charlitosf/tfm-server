package controllers

import (
	"charlitosf/tfm-server/internal/services"
	"charlitosf/tfm-server/pkg/httptypes"

	"github.com/gin-gonic/gin"
)

// Get all files handler
func GetFiles(c *gin.Context) {
	// Get user from the context
	user := c.MustGet("username").(string)
	// Call services GetAllFilenames method
	files, err := services.GetAllFilenames(user)
	if err == nil {
		c.JSON(200, files)
	} else {
		c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
	}
}

// Create a file handler
func CreateFile(c *gin.Context) {
	// Get user from the context
	user := c.MustGet("username").(string)
	// Bind request body to CreateFile struct
	var createFile httptypes.CreateFileRequest
	err := c.BindJSON(&createFile)
	if err == nil {
		// Call services CreateFile method
		err = services.CreateFile(user, createFile.Name, createFile.Content)
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
	// Get user from the context
	user := c.MustGet("username").(string)
	// Get filename from the path
	filename := c.Param("name")
	// Call services GetFile method
	data, err := services.GetFile(user, filename)
	if err == nil {
		c.JSON(200, httptypes.GetFileResponse{
			Name:    filename,
			Content: data,
		})
	} else {
		c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
	}
}

// Update a file handler
func UpdateFile(c *gin.Context) {
	// Get user from the context
	user := c.MustGet("username").(string)
	// Get filename from the path
	filename := c.Param("name")
	// Bind request body to UpdateFile struct
	var updateFile httptypes.UpdateFileRequest
	err := c.BindJSON(&updateFile)
	if err == nil {
		// Call services UpdateFile method
		err = services.UpdateFile(user, filename, updateFile.Content)
		if err == nil {
			c.JSON(200, httptypes.GenericResponse{})
		} else {
			c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
		}
	} else {
		c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
	}
}

// Delete a file handler
func DeleteFile(c *gin.Context) {
	// Get user from the context
	user := c.MustGet("username").(string)
	// Get filename from the path
	filename := c.Param("name")
	// Call services DeleteFile method
	err := services.DeleteFile(user, filename)
	if err == nil {
		c.JSON(200, httptypes.GenericResponse{})
	} else {
		c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
	}
}
