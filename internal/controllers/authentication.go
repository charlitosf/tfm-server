package controllers

import (
	"charlitosf/tfm-server/internal/services"
	"charlitosf/tfm-server/pkg/httptypes"
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
	var req httptypes.SignupRequest
	err := c.BindJSON(&req)
	if err == nil { // Correct request
		// Perform signup
		err := services.Signup(req.Username, req.Password)
		if err != nil { // Username already exists or other error
			c.JSON(400, httptypes.SignupResponse{Error: &httptypes.Error{Message: err.Error()}})
		} else {
			c.JSON(200, httptypes.SignupResponse{})
		}
	} else {
		// If the request is incorrect, abort with an error
		c.AbortWithError(400, err)
	}
}
