package model

// ChampionDataExtended contains additional data about a champion
type ChampionDataExtended struct {
	ChampionData
	Skins            []SkinData            `json:"skins"`
	Lore             string                `json:"lore"`
	AllyTips         []string              `json:"allytips"`
	EnemyTips        []string              `json:"enemytips"`
	Spells           []SpellData           `json:"spells"`
	Passive          PassiveData           `json:"passive"`
	RecommendedItems []RecommendedItemData `json:"recommended"`
}
