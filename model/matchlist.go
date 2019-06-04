package model

// Matchlist contains information about all games played by a single summoner
type Matchlist struct {
	Matches    []MatchReference `json:"matches"`
	TotalGames int              `json:"totalGames"`
	StartIndex int              `json:"startIndex"`
	EndIndex   int              `json:"endIndex"`
}
