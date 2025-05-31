package dbhandler

import (
	"github.com/inonsdn/gacha-system/gacha_service/internal/constants"
)

type DBHandler struct {
	// Mocking
	GachaHistory []map[string]string
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
	return map[string]int{
		constants.CommonRarity:    100,
		constants.RareRarity:      50,
		constants.SuperRareRarity: 20,
		constants.LegendRarity:    10,
	}
}

func NewDBHandler() *DBHandler {
	return &DBHandler{
		GachaHistory: []map[string]string{},
	}
}
