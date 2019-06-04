package model

// Perks represents the runes for a player in an ongoing game
type Perks struct {
	PerkStyle    int   `json:"perkStyle"`
	PerksIDs     []int `json:"perksIDs"`
	PerkSubStyle int   `json:"perkSubStyle"`
}
