package servicemanager

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/inonsdn/gacha-system/api_gateway/internal/client"
	"github.com/inonsdn/gacha-system/api_gateway/internal/middleware"
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

		// get param from request path
		gachaId := c.Param("gachaId")
		drawAmount := c.Query("amount")

		if drawAmount == "" {
			drawAmount = "1"
		}

		drawAmountInt, err := strconv.Atoi(drawAmount)
		if err != nil {
			fmt.Println("Invalid number:", err)
			return
		}

		// Call GachaService via gRPC
		result, err := gachaClient.Draw(userID, gachaId, int32(drawAmountInt))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gacha draw failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"items": result.Items,
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
			"gachaCateg":  result.Name,
			"remaining":   result.RarityInfos,
			"startDate":   result.StartDate,
			"endDate":     result.EndDate,
			"ownerUserId": result.OwnerUserId,
		})
	}
}
