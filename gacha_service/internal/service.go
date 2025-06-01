package internal

import (
	"context"
	"fmt"

	"github.com/inonsdn/gacha-system/gacha_service/internal/dbhandler"
	gachapb "github.com/inonsdn/gacha-system/proto/gacha"
)

type GachaService struct {
	DBHandler *dbhandler.DBHandler
	gachapb.UnimplementedGachaServiceServer
}

type GachaItem struct {
	index  int
	name   string
	rarity string
	categ  string
}

type GacheCategInfo struct {
	name              string
	rarityToRemaining map[string]int
}

func buildDrawResponse(gachaItems []GachaItem) *gachapb.DrawResponse {
	var res []*gachapb.DrawItem

	for _, item := range gachaItems {
		res = append(res, &gachapb.DrawItem{
			Index:  fmt.Sprintf("%d", item.index),
			Name:   item.name,
			Rarity: item.rarity,
		})
	}

	return &gachapb.DrawResponse{
		Items: res,
	}
}

func buildGachaResponse(gacheCategInfo GacheCategInfo) *gachapb.GachaResponse {
	return &gachapb.GachaResponse{
		Name: gacheCategInfo.name,
		RarityInfos: []*gachapb.GacheRarityRemain{
			{
				Rarity:    "SR",
				Remaining: 9999,
			},
		},
		StartDate:   "20250501",
		EndDate:     "20300501",
		OwnerUserId: "1",
	}
}

func (g GachaService) Draw(c context.Context, drawRequest *gachapb.DrawRequest) (*gachapb.DrawResponse, error) {
	// TODO: implement logic to draw gacha
	var err error

	fmt.Println("Draw called", drawRequest)

	gachaItems := normalDraw(g.DBHandler, drawRequest.UserId, drawRequest.GachaId, drawRequest.DrawAmount)

	// construct gacha response
	drawResponse := buildDrawResponse(gachaItems)
	fmt.Println("Draw res", drawResponse)

	return drawResponse, err
}

func (g GachaService) GetGachaInfo(c context.Context, gachaRequest *gachapb.GachaRequest) (*gachapb.GachaResponse, error) {
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
