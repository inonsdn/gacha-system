package dbhandler

import (
	"fmt"

	"github.com/inonsdn/gacha-system/gacha_service/internal/constants"

	"github.com/google/uuid"
)

type DBHandler struct {
	// Mocking
	GachaHistory []map[string]string

	GachaIdToRemaining map[string]map[string]int
}

type QueryItem struct {
}

type QueryResult struct {
	items []QueryItem
}

func (db *DBHandler) QueryFromLocal() {

}

func (db *DBHandler) GetGachaHistory() []map[string]string {
	return db.GachaHistory
}

func (db *DBHandler) GetGachaRemaining(categId string) map[string]int {
	return db.GachaIdToRemaining[categId]
}

func NewDBHandler() *DBHandler {
	gachaId1 := uuid.New().String()
	gachaId2 := uuid.New().String()
	fmt.Println("===== Gacha ID 1", gachaId1)
	fmt.Println("===== Gacha ID 2", gachaId2)
	return &DBHandler{
		GachaHistory: []map[string]string{},
		GachaIdToRemaining: map[string]map[string]int{
			gachaId1: {
				constants.CommonRarity:    100,
				constants.RareRarity:      50,
				constants.SuperRareRarity: 20,
				constants.LegendRarity:    10,
			},
			gachaId2: {
				constants.CommonRarity:    100,
				constants.RareRarity:      50,
				constants.SuperRareRarity: 20,
				constants.LegendRarity:    10,
			},
		},
	}
}
