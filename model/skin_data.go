package model

// SkinData contains information about a skin for a champion
type SkinData struct {
	ID      string `json:"id"`
	Num     int    `json:"num"`
	Name    string `json:"name"`
	Chromas bool   `json:"chromas"`
}
