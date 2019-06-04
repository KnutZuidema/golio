package model

// RecommendedItem represents an item in a recommended set
type RecommendedItem struct {
	ID        string `json:"id"`
	Count     int    `json:"count"`
	HideCount bool   `json:"hideCount"`
}
