package model

// ParticipantTimeline contains timeline values for a participant in a game
type ParticipantTimeline struct {
	// Participant's calculated lane. MID and BOT are legacy values.
	// (Legal values: MID, MIDDLE, TOP, JUNGLE, BOT, BOTTOM)
	Lane          string `json:"lane"`
	ParticipantID int    `json:"participantId"`
	// Creep score difference versus the calculated lane opponent(s) for a specified period.
	CsDiffPerMinDeltas map[string]float64 `json:"csDiffPerMinDeltas"`
	// Gold for a specified period.
	GoldPerMinDeltas map[string]float64 `json:"goldPerMinDeltas"`
	// Experience difference versus the calculated lane opponent(s) for a specified period.
	XpDiffPerMinDeltas map[string]float64 `json:"xpDiffPerMinDeltas"`
	// Creeps for a specified period.
	CreepsPerMinDeltas map[string]float64 `json:"creepsPerMinDeltas"`
	// Experience change for a specified period.
	XpPerMinDeltas map[string]float64 `json:"xpPerMinDeltas"`
	// Participant's calculated role. (Legal values: DUO, NONE, SOLO, DUO_CARRY, DUO_SUPPORT)
	Role string `json:"role"`
	// Damage taken difference versus the calculated lane opponent(s) for a specified period.
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	// Damage taken for a specified period.
	DamageTakenPerMinDeltas map[string]float64 `json:"damageTakenPerMinDeltas"`
}
