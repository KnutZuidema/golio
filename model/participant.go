package model

// Participant represents a participant in a game
type Participant struct {
	// Participant statistics.
	Stats         ParticipantStats `json:"stats"`
	ParticipantID int              `json:"participantId"`
	// List of legacy Rune information. Not included for matches played with Runes Reforged.
	Runes []Rune `json:"runes"`
	// Participant timeline data.
	Timeline ParticipantTimeline `json:"timeline"`
	// 100 for blue side. 200 for red side.
	TeamID int `json:"teamId"`
	// Second Summoner Spell id.
	Spell2ID int `json:"spell2Id"`
	// List of legacy Mastery information. Not included for matches played with Runes Reforged.
	Masteries []LegacyMastery `json:"masteries"`
	// Highest ranked tier achieved for the previous season in a specific subset of queueIds, if any, otherwise null.
	// Used to display border in game loading screen. Please refer to the Ranked Info documentation.
	// (Legal values: CHALLENGER, MASTER, DIAMOND, PLATINUM, GOLD, SILVER, BRONZE, UNRANKED)
	HighestAchievedSeasonTier string `json:"highestAchievedSeasonTier"`
	// First Summoner Spell id.
	Spell1ID   int `json:"spell1Id"`
	ChampionID int `json:"championId"`
}
