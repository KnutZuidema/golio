package model

// MatchReference contains information about a game by a single summoner
type MatchReference struct {
	Lane       string `json:"lane"`
	GameID     int    `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Season     int    `json:"season"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Timestamp  int    `json:"timestamp"`
}
