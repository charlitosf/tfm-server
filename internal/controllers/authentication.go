package controllers

import (
	"charlitosf/tfm-server/internal/services"
	"charlitosf/tfm-server/pkg/httptypes"

	"github.com/gin-gonic/gin"
)

// Login handler
func Login(c *gin.Context) {
	// Get the user credentials
	var req httptypes.LoginRequest
	err := c.BindJSON(&req)
	if err == nil { // Correct request
		// Get the token and the user
		token, user, err := services.Login(req.Username, req.Password)
		if err != nil {
			c.JSON(400, httptypes.LoginResponse{Error: &httptypes.Error{Message: err.Error()}})
		} else {
			// Return the token and the user
			c.JSON(200, httptypes.LoginResponse{Token: *token, User: &httptypes.UserMetadata{
				Name:    user.Name,
				Email:   user.Email,
				PubKey:  user.PubKey,
				PrivKey: user.PrivKey,
			}})
		}
	} else { // Wrong request
		c.JSON(400, httptypes.LoginResponse{Error: &httptypes.Error{Message: err.Error()}})
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
		err := services.Signup(req.Username, req.Password, req.Name, req.Email, req.PubKey, req.PrivKey)
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
