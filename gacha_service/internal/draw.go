package internal

import (
	"math/rand"
	"time"

	"github.com/inonsdn/gacha-system/gacha_service/internal/constants"
	"github.com/inonsdn/gacha-system/gacha_service/internal/dbhandler"
)

func generatePossibleList(rareAmount int, superRareAmount int, LegendAmount int) []string {
	var possibleList []string
	for i := 0; i < rareAmount; i++ {
		possibleList = append(possibleList, constants.RareRarity)
	}

	for i := 0; i < superRareAmount; i++ {
		possibleList = append(possibleList, constants.SuperRareRarity)
	}

	for i := 0; i < LegendAmount; i++ {
		possibleList = append(possibleList, constants.LegendRarity)
	}

	return possibleList
}

func computeProability() {

}

func normalDraw(dbHandler *dbhandler.DBHandler, userId string, gachaId string, drawAmount int) []GachaItem {

	var items []GachaItem

	rarityToRemaining := dbHandler.GetGachaRemaining("test01")
	possibleList := generatePossibleList(rarityToRemaining[constants.RareRarity], rarityToRemaining[constants.SuperRareRarity], rarityToRemaining[constants.LegendRarity])

	rand.NewSource(time.Now().UnixNano())
	for index := range drawAmount {
		randomNum := rand.Intn(len(possibleList))
		randomRarity := possibleList[randomNum]

		items = append(items, GachaItem{
			index:  index,
			name:   "Test",
			rarity: randomRarity,
			categ:  gachaId,
		})
	}

	return items
}

// func drawGacha(dbHandler *dbhandler.DBHandler, userId string, gachaId string, drawAmount int) []GachaItem {
// 	// get history of gacha and user
// 	histories := dbHandler.GetGachaHistory()

// 	var items []GachaItem

// 	for index := range drawAmount {
// 		items = append(items, GachaItem{
// 			index:  index,
// 			name:   "Test",
// 			rarity: constants.CommonRarity,
// 			categ:  gachaId,
// 		})
// 	}

// 	return items
// }
