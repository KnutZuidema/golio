package model

// CurrentGameParticipant represents a player in an ongoing game
type CurrentGameParticipant struct {
	ProfileIconID            int                       `json:"profileIconId"`
	ChampionID               int                       `json:"championId"`
	SummonerName             string                    `json:"summonerName"`
	GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
	Bot                      bool                      `json:"bot"`
	Perks                    Perks                     `json:"perks"`
	Spell2ID                 int                       `json:"spell2Id"`
	TeamID                   int                       `json:"teamId"`
	Spell1ID                 int                       `json:"spell1Id"`
	SummonerID               string                    `json:"summonerId"`
}
