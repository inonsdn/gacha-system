package servicemanager

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inonsdn/gacha-system/api_gateway/internal/client"
	"github.com/inonsdn/gacha-system/api_gateway/internal/middleware"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Login(userClient client.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("API GOT LOGIN")
		loginName, _ := c.Params.Get("loginName")
		passwd, _ := c.Params.Get("passwd")
		// call function to user service via grpc
		userLoginResponse, err := userClient.Login(loginName, passwd)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Cannot login.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"token": userLoginResponse.JwtToken,
			})
		}

	}
}

func Register(userClient client.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		// call function to user service via grpc
		userClient.Register("", "")

		c.JSON(http.StatusOK, gin.H{
			"message": "Done",
		})
	}
}

type DrawRequestBody struct {
	GachaId string `json:"gachaId"`
	Amount  int32  `json:"amount"`
}

func GachaDraw(gachaClient client.GachaServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("GACHADRAW: called")
		userID := middleware.GetUserIDFromContext(c)

		// Parse payload to body request object
		var body DrawRequestBody
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			return
		}

		// Call GachaService via gRPC
		result, err := gachaClient.Draw(userID, body.GachaId, body.Amount)
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
