package model

// TournamentCodeParameters parameters needed to create tournament codes
type TournamentCodeParameters struct {
	// The spectator type of the game. (Legal values: NONE, LOBBYONLY, ALL)
	SpectatorType string `json:"spectatorType"`
	// The team size of the game. Valid values are 1-5.
	TeamSize int `json:"teamSize"`
	// The pick type of the game. (Legal values: BLIND_PICK, DRAFT_MODE, ALL_RANDOM, TOURNAMENT_DRAFT)
	PickType string `json:"pickType"`
	// Optional list of encrypted summonerIds in order to validate the players eligible to join the lobby.
	// NOTE: We currently do not enforce participants at the team level, but rather the aggregate of teamOne and
	// teamTwo. We may add the ability to enforce at the team level in the future.
	AllowedSummonerIDs []string `json:"allowedSummonerIds"`
	// The map type of the game. (Legal values: SUMMONERS_RIFT, TWISTED_TREELINE, HOWLING_ABYSS)
	MapType string `json:"mapType"`
	// Optional string that may contain any data in any format, if specified at all. Used to denote any custom
	// information about the game.
	Metadata string `json:"metadata"`
}
