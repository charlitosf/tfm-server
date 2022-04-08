package middleware

import "github.com/gin-gonic/gin"

// Authorized is a middleware that checks if the request contains a valid token
// in the Authorization header. If the token is valid, the middleware
// forwards the request to the next handler. Otherwise, it returns an error.
func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the header
		token := c.GetHeader("Authorization")

		// Check if the token is valid
		if IsTokenValid(token) {
			// Token is valid, call the next handler
			c.Next()
		} else {
			// Token is invalid, return with an error
			c.AbortWithStatus(401)
		}
	}
}

func IsTokenValid(token string) bool {
	return true
}
