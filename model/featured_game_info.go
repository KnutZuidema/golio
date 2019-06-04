package model

type FeaturedGameInfo struct {
	GameID            int              `json:"gameId"`
	GameStartTime     int              `json:"gameStartTime"`
	PlatformID        string           `json:"platformId"`
	GameMode          string           `json:"gameMode"`
	MapID             int              `json:"mapId"`
	GameType          string           `json:"gameType"`
	BannedChampions   []BannedChampion `json:"bannedChampions"`
	Observers         Observer         `json:"observers"`
	Participants      []Participant    `json:"participants"`
	GameLength        int              `json:"gameLength"`
	GameQueueConfigID int              `json:"gameQueueConfigId"`
}
