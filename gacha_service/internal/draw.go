package internal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inonsdn/gacha-system/gacha_service/internal/constants"
	"github.com/inonsdn/gacha-system/gacha_service/internal/dbhandler"
)

func computeProability() {

}

func normalDraw(dbHandler *dbhandler.DBHandler, userId string, gachaId string, drawAmount int32) []GachaItem {

	var items []GachaItem

	rarityToRemaining := dbHandler.GetGachaRemaining(gachaId)

	if len(rarityToRemaining) == 0 {
		fmt.Println("Not found gacha id: ", gachaId)
		return items
	}

	var randomRarity string
	rand.NewSource(time.Now().UnixNano())
	for index := range int(drawAmount) {

		// random num
		allItemsNum := rarityToRemaining[constants.CommonRarity] + rarityToRemaining[constants.RareRarity] + rarityToRemaining[constants.SuperRareRarity] + rarityToRemaining[constants.LegendRarity]
		randomNum := rand.Intn(allItemsNum)

		// categorize random num to be rarity of item
		if randomNum < rarityToRemaining[constants.CommonRarity] {
			randomRarity = constants.CommonRarity
		} else if randomNum < rarityToRemaining[constants.CommonRarity]+rarityToRemaining[constants.RareRarity] {
			randomRarity = constants.RareRarity
		} else if randomNum < rarityToRemaining[constants.CommonRarity]+rarityToRemaining[constants.RareRarity]+rarityToRemaining[constants.SuperRareRarity] {
			randomRarity = constants.SuperRareRarity
		} else {
			randomRarity = constants.LegendRarity
		}

		// update remaining item
		dbHandler.GachaIdToRemaining[gachaId][randomRarity] = dbHandler.GachaIdToRemaining[gachaId][randomRarity] - 1

		items = append(items, GachaItem{
			index:  index,
			name:   "Test",
			rarity: randomRarity,
			categ:  gachaId,
		})
	}
	rarityToRemaining = dbHandler.GetGachaRemaining(gachaId)
	fmt.Println("========== rarityToRemaining: ", rarityToRemaining)

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
