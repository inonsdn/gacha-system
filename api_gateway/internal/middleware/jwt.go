package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthJWT - Middleware to check JWT token
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		// token := c.GetHeader("Authorization")
		// if token == "" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Token required"})
		// 	c.Abort()
		// 	return
		// }

		// // Validate token (for simplicity, we assume it's always valid here)
		// if !validateToken(token) {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		// 	c.Abort()
		// 	return
		// }

		// Assuming the token contains the userID, we set it on the context
		c.Set("userID", "12345") // For example, this should be extracted from the JWT

		// Proceed to the next handler
		c.Next()
	}
}

// validateToken - A basic example function to validate the token (you can replace with actual JWT validation)
func validateToken(token string) bool {
	// In a real-world scenario, decode and validate JWT here
	// TODO:
	return strings.HasPrefix(token, "Bearer ")
}

// GetUserIDFromContext - Utility to extract userID from context
func GetUserIDFromContext(c *gin.Context) string {
	userID, _ := c.Get("userID")
	return userID.(string)
}
