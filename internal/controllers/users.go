package controllers

import (
	"charlitosf/tfm-server/internal/services"
	"charlitosf/tfm-server/pkg/httptypes"

	"github.com/gin-gonic/gin"
)

// Update user ('s password) handler
func UpdateUser(c *gin.Context) {
	// Bind request body
	var request httptypes.UpdateUserRequest
	err := c.BindJSON(&request)
	if err == nil {
		// Update user
		err = services.UpdateUserPassword(c.MustGet("username").(string), request.Password, c.MustGet("token").(string), c.MustGet("xtotp").(string))
		if err != nil {
			c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
		} else {
			c.JSON(200, httptypes.GenericResponse{})
		}
	} else {
		c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
	}
}

// Delete user handler
func DeleteUser(c *gin.Context) {
	err := services.DeleteUser(c.Param("username"), c.MustGet("token").(string), c.MustGet("xtotp").(string))
	if err != nil {
		c.JSON(400, httptypes.GenericResponse{Error: &httptypes.Error{Message: err.Error()}})
	} else {
		c.JSON(200, httptypes.GenericResponse{})
	}
}
