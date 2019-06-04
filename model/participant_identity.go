package model

// ParticipantIdentity contains a reference to a player for a participant in a game
type ParticipantIdentity struct {
	// Player information.
	Player        Player `json:"player"`
	ParticipantID int    `json:"participantId"`
}
