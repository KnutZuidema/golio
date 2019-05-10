package model

type ChampionInfo struct {
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIDsForNewPlayers"`
	FreeChampionIDs              []int `json:"freeChampionIDs"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}
