package internal

import (
	"context"
	"fmt"

	"local.dev/gacha_service/proto_file/gacha"
)

type GachaService struct {
	gacha.UnimplementedGachaServiceServer
}

type GachaItem struct {
	name   string
	rarity string
	categ  string
}

type GacheCategInfo struct {
	name              string
	rarityToRemaining map[string]int
}

func buildDrawResponse(gachaItem GachaItem) *gacha.DrawResponse {
	return &gacha.DrawResponse{
		ItemName: gachaItem.name,
		Rarity:   gachaItem.rarity,
	}
}

func buildGachaResponse(gacheCategInfo GacheCategInfo) *gacha.GachaResponse {
	return &gacha.GachaResponse{
		GachaCateg:  gacheCategInfo.name,
		Remaining:   "0",
		StartDate:   "20250501",
		EndDate:     "20300501",
		OwnerUserId: "1",
	}
}

func (g GachaService) Draw(c context.Context, drawRequest *gacha.DrawRequest) (*gacha.DrawResponse, error) {
	// TODO: implement logic to draw gacha
	var err error

	fmt.Println("Draw called", drawRequest)
	// Mocking result of draw, which is gacha item
	gachaItem := GachaItem{
		name:   "Test_0000",
		rarity: "SSR",
		categ:  "TestNo0",
	}

	// construct gacha response
	drawResponse := buildDrawResponse(gachaItem)
	fmt.Println("Draw res", drawResponse)

	return drawResponse, err
}

func (g GachaService) GetGachaInfo(c context.Context, gachaRequest *gacha.GachaRequest) (*gacha.GachaResponse, error) {
	var err error
	fmt.Println("GetGachaInfo called", gachaRequest)
	// query gacha info from db

	// Mocking result from db
	gachaInfo := GacheCategInfo{
		name: "TestNo0",
		rarityToRemaining: map[string]int{
			"R":    9999,
			"SR":   9999,
			"SSR":  9999,
			"USSR": 9999,
		},
	}

	gachaResponse := buildGachaResponse(gachaInfo)
	fmt.Println("GetGachaInfo res", gachaResponse)

	return gachaResponse, err
}
