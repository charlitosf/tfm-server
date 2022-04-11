package main

import (
	"charlitosf/tfm-server/internal/controllers"
	"charlitosf/tfm-server/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Main router object
	r := gin.Default()

	// Routes

	// Authorization group
	auth := r.Group("/auth")
	{
		// Login
		auth.POST("/login", controllers.Login)
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
		users := authorized.Group("/users")
		{
			// Update user ('s password)
			users.PUT("/:id", controllers.UpdateUser)
			// Delete user
			users.DELETE("/:id", controllers.DeleteUser)
		}

		// Passwords group
		passwords := authorized.Group("/passwords")
		{
			// Get all passwords
			passwords.GET("/", controllers.GetPasswords)
			// Get all passwords of a website
			passwords.GET("/:website", controllers.GetPasswordsByWebsite)
			// Create a password on a website
			passwords.POST("/:website", controllers.CreatePassword)
			// Get a password of a website of a user
			passwords.GET("/:website/:username", controllers.GetPassword)
			// Update a password of a website of a user
			passwords.PUT("/:website/:username", controllers.UpdatePassword)
			// Delete a password of a website of a user
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

	// Start the server
	r.Run()
}
