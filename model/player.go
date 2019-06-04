package model

// Player represents a player
type Player struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	// Original platformId.
	PlatformID string `json:"platformId"`
	// Player's current accountID (Encrypted)
	CurrentAccountID string `json:"currentAccountId"`
	ProfileIcon      int    `json:"profileIcon"`
	// Player's summonerID (Encrypted)
	SummonerID string `json:"summonerId"`
	// Player's original accountID (Encrypted)
	AccountID string `json:"accountId"`
}
