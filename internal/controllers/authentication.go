package controllers

import (
	"charlitosf/tfm-server/internal/services"
	"errors"

	"github.com/gin-gonic/gin"
)

// Login handler
func Login(c *gin.Context) {
	// Get the user credentials
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Check credentials
	if username == "admin" && password == "admin" {
		// Create token
		//token := CreateToken(username)
		token := 0
		// Set the token in the response
		c.JSON(200, gin.H{
			"token": token,
		})
	} else {
		// If the credentials are wrong, abort with an error
		c.AbortWithError(401, errors.New("wrong credentials"))
	}
}

// Logout handler
func Logout(c *gin.Context) {
	// Delete the token from the database
	// DeleteToken(c.GetHeader("Authorization"))
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// Signup handler
func Signup(c *gin.Context) {
	// Perform signup
	err := services.Signup("admin", "admin")
	c.JSON(200, gin.H{
		"status": err.Error(),
	})
}
