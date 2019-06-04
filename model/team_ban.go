package model

// TeamBan is a champion banned by a team
type TeamBan struct {
	// Turn during which the champion was banned.
	PickTurn int `json:"pickTurn"`
	// Banned championId.
	ChampionID int `json:"championId"`
}
