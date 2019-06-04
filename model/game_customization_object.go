package model

// GameCustomizationObject contains information specific to an ongoing game
type GameCustomizationObject struct {
	Category string `json:"category"`
	Content  string `json:"content"`
}
