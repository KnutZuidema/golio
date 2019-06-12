package model

// LobbyEvent represents an event that happened in a tournament lobby
type LobbyEvent struct {
	EventType  string `json:"eventType"`
	SummonerID string `json:"summonerId"`
	Timestamp  string `json:"timestamp"`
}
