package controllers

import (
	"charlitosf/tfm-server/internal/services"
	"charlitosf/tfm-server/pkg/httptypes"

	"github.com/gin-gonic/gin"
)

// Update user ('s password) handler
func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Delete user handler
func DeleteUser(c *gin.Context) {
	err := services.DeleteUser(c.Param("username"), c.MustGet("token").(string))
	if err != nil {
		c.JSON(400, httptypes.DeleteUserResponse{Error: &httptypes.Error{Message: err.Error()}})
	} else {
		c.JSON(200, httptypes.DeleteUserResponse{})
	}
}
