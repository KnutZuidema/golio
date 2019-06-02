package model

// ChampionMastery represents the mastery of a champion in the mastery system for a summoner
type ChampionMastery struct {
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionLevel                int    `json:"championLevel"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionID                   int    `json:"championId"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	LastPlayTime                 int    `json:"lastPlayTime"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	SummonerID                   string `json:"summonerId"`
}
