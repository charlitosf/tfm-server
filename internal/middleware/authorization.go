package middleware

import (
	"charlitosf/tfm-server/internal/jwt"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authorized is a middleware that checks if the request contains a valid token
// in the Authorization header. If the token is valid, the middleware
// forwards the request to the next handler. Otherwise, it returns an error.
// Gin-gonic authorization middleware function
func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authHeader from the header
		authHeader := c.Request.Header.Get("Authorization")

		// Validate the token
		if strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")
			// Add the token to the context
			c.Set("token", token)

			// Validate the token
			username, err := jwt.VerifyJWT(token)
			if err != nil {
				// If token is invalid, abort with error
				c.AbortWithError(401, errors.New("unauthorized token"))
			} else {
				// If token is valid, set the username to the context
				c.Set("username", username)
			}
		} else {
			// If token is missing, abort with error
			c.AbortWithError(401, errors.New("missing or invalid token"))
		}
	}
}

// Token username and path username must match
func TokenUsernameMustMatchPathUsername() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the username from the context
		username := c.MustGet("username").(string)

		// Get the path username from the path
		pathUsername := c.Param("username")

		// If the username from the path does not match the username from the context, abort with error
		if username != pathUsername {
			c.AbortWithError(401, errors.New("token username does not match path username"))
		}
	}
}

// Add CORS headers to the response
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Totp")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
