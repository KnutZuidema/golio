package model

// TournamentUpdateParameters parameters needed to update an existing tournament
type TournamentUpdateParameters struct {
	// The spectator type (Legal values: NONE, LOBBYONLY, ALL)
	SpectatorType string `json:"spectatorType"`
	// The pick type (Legal values: BLIND_PICK, DRAFT_MODE, ALL_RANDOM, TOURNAMENT_DRAFT)
	PickType string `json:"pickType"`
	// Optional list of encrypted summonerIds in order to validate the players eligible to join the lobby.
	// NOTE: Participants are not enforced at the team level, but rather the aggregate of teamOne and teamTwo.
	AllowedSummonerIDs []string `json:"allowedSummonerIds"`
	// The map type (Legal values: SUMMONERS_RIFT, TWISTED_TREELINE, HOWLING_ABYSS)
	MapType string `json:"mapType"`
}
