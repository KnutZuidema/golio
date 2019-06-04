package model

// RecommendedItemData a build recommended for a champion
type RecommendedItemData struct {
	Champion      string               `json:"champion"`
	Title         string               `json:"title"`
	Map           string               `json:"map"`
	Mode          string               `json:"mode"`
	CustomTag     string               `json:"customTag"`
	SortRank      int                  `json:"sortrank"`
	ExtensionPage bool                 `json:"extensionPage"`
	CustomPanel   interface{}          `json:"customPanel"`
	Blocks        []RecommendedItemSet `json:"blocks"`
}
