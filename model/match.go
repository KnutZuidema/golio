package model

// Match contains information about a match
type Match struct {
	// Please refer to the Game Constants documentation.
	SeasonID int `json:"seasonId"`
	// Please refer to the Game Constants documentation.
	QueueID int `json:"queueId"`
	GameID  int `json:"gameId"`
	// Participant identity information.
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	// The major.minor version typically indicates the patch the match was played on.
	GameVersion string `json:"gameVersion"`
	// Platform where the match was played.
	PlatformID string `json:"platformId"`
	// Please refer to the Game Constants documentation.
	GameMode string `json:"gameMode"`
	// Please refer to the Game Constants documentation.
	MapID int `json:"mapId"`
	// Please refer to the Game Constants documentation.
	GameType string `json:"gameType"`
	// Team information.
	Teams []TeamStats `json:"teams"`
	// Participant information.
	Participants []Participant `json:"participants"`
	// Match duration in seconds.
	GameDuration int `json:"gameDuration"`
	// Designates the timestamp when champion select ended and the loading screen appeared, NOT when the game timer was
	// at 0:00.
	GameCreation int `json:"gameCreation"`
}
