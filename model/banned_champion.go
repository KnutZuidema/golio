package model

type BannedChampion struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
	TeamID     int `json:"teamId"`
}
