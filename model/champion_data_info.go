package model

// ChampionDataInfo contains information about the playstyle of a champion
type ChampionDataInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}
