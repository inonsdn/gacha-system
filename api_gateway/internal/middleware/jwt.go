package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("ebcc8244-5629-4633-bbdd-f5f2253a13bd") // 🔐 Use env var in real apps

// AuthJWT - Middleware to check JWT token
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// read authroize header
		authHeader := c.GetHeader("Authorization")

		fmt.Println("[Authorization] has request")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			return
		}

		// extract only token from header
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		}, jwt.WithValidMethods([]string{"HS256"}))

		if err != nil || !token.Valid {
			fmt.Println("[Authorization] Invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			fmt.Println("[Authorization] Cannot get user id")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}

		// expirationTime, err := token.Claims.GetExpirationTime()
		// if expirationTime.Before(time.Now()) {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expire, login again"})
		// }

		userID := claims["user_id"].(string)
		c.Set("userID", userID) // ✅ Save for handler

		fmt.Println("[Authorization] Authorized. User: ", userID)

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
