package middleware

import (
	"charlitosf/tfm-server/pkg/stringutilities"

	"github.com/gin-gonic/gin"
)

// Middleware that reverses the website from the path
// using the function of the string utilities package
func ReverseWebsite() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the website from the path
		website := c.Param("website")

		// Reverse the website
		website = stringutilities.ReverseSplitJoin(website)

		// Set the website to the context
		c.Set("website", website)
	}
}
