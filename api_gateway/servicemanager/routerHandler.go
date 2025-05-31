package servicemanager

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"local.dev/api_gateway/internal/client"
	"local.dev/api_gateway/internal/middleware"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Redirect to login",
	})
}

func GachaDraw(gachaClient client.GachaServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserIDFromContext(c)

		// Call GachaService via gRPC
		result, err := gachaClient.Draw(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gacha draw failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"item":   result.ItemName,
			"rarity": result.Rarity,
		})
	}
}

func GetGachaInfo(gachaClient client.GachaServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserIDFromContext(c)

		// get param from request path
		gachaCateg := c.Param("categ")

		// Call GachaService via gRPC
		result, err := gachaClient.GetGachaInfo(userID, gachaCateg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gacha info failed", "err": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"gachaCateg":  result.GachaCateg,
			"remaining":   result.Remaining,
			"startDate":   result.StartDate,
			"endDate":     result.EndDate,
			"ownerUserId": result.OwnerUserId,
		})
	}
}
