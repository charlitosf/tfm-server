package main

import (
	"charlitosf/tfm-server/internal/controllers"
	"charlitosf/tfm-server/internal/crypt"
	"charlitosf/tfm-server/internal/middleware"
	"charlitosf/tfm-server/internal/services"
	"fmt"
	"syscall"

	"github.com/gin-gonic/gin"
	"golang.org/x/term"
)

func main() {
	pass, err := getServerPassword()
	if err != nil {
		panic(err)
	}
	services.ServerPassword = crypt.Hash256(pass)

	// Main router object
	r := gin.Default()
	r.Use(middleware.CORS())
	// Routes

	// Authorization group
	auth := r.Group("/auth")
	{
		// Login (must contain the X-TOTP header)
		auth.POST("/login", middleware.XTOTP(), controllers.Login)
		// Logout (must be authorized)
		auth.POST("/logout", middleware.Authorized(), controllers.Logout)
		// Signup
		auth.POST("/signup", controllers.Signup)
	}

	// Rest of the routes that must be authorized
	authorized := r.Group("/")

	// Middleware
	authorized.Use(middleware.Authorized())
	{
		// Users group
		users := authorized.Group("/users", middleware.TokenUsernameMustMatchPathUsername())
		{
			// Update user ('s password) (must contain the X-TOTP header)
			users.PUT("/:username", middleware.XTOTP(), controllers.UpdateUser)
			// Delete user ('s account)	(must contain the X-TOTP header)
			users.DELETE("/:username", middleware.XTOTP(), controllers.DeleteUser)
		}

		// Passwords group
		passwords := authorized.Group("/passwords")
		{
			// Get all passwords (must contain the X-TOTP header)
			passwords.GET("/", middleware.XTOTP(), controllers.GetPasswords)
			// Get all passwords from a website
			passwords.GET("/:website", controllers.GetPasswordsByWebsite)
			// Create a password on a website
			passwords.POST("/:website", controllers.CreatePassword)
			// Get a password from a website of a user
			passwords.GET("/:website/:username", controllers.GetPassword)
			// Update a password from a website of a user
			passwords.PUT("/:website/:username", controllers.UpdatePassword)
			// Delete a password from a website of a user
			passwords.DELETE("/:website/:username", controllers.DeletePassword)
		}

		// Files group
		files := authorized.Group("/files")
		{
			// Get all files
			files.GET("/", controllers.GetFiles)
			// Create a file
			files.POST("/", controllers.CreateFile)
			// Get a file
			files.GET("/:name", controllers.GetFile)
			// Update a file
			files.PUT("/:name", controllers.UpdateFile)
			// Delete a file
			files.DELETE("/:name", controllers.DeleteFile)
		}
	}

	// Start the server with TLS
	r.RunTLS(":8080", "192.168.22.132.crt", "192.168.22.132.key")
}

// Function that reads the server password from the terminal
func getServerPassword() (string, error) {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	return string(bytePassword), nil
}
