package model

// LobbyEventList is a wrapper for a list of lobby events in a tournament
type LobbyEventList struct {
	EventList []LobbyEvent `json:"eventList"`
}
