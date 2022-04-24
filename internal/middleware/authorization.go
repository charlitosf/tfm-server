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
