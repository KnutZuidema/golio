package model

// BannedChampion represents a champion ban during pack/ban phase
type BannedChampion struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
	TeamID     int `json:"teamId"`
}
