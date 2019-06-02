package model

// ChampionInfo contains information about the free champion rotation
type ChampionInfo struct {
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIDsForNewPlayers"`
	FreeChampionIDs              []int `json:"freeChampionIDs"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}
