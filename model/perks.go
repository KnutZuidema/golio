package model

type Perks struct {
	PerkStyle    int   `json:"perkStyle"`
	PerksIDs     []int `json:"perksIDs"`
	PerkSubStyle int   `json:"perkSubStyle"`
}
