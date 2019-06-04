package model

// LeagueEntry represents a summoners ranked position in a league
type LeagueEntry struct {
	Queue        string     `json:"queueType"`
	SummonerName string     `json:"summonerName"`
	HotStreak    bool       `json:"hotStreak"`
	MiniSeries   MiniSeries `json:"miniSeries"`
	Wins         int        `json:"wins"`
	Veteran      bool       `json:"veteran"`
	Losses       int        `json:"losses"`
	FreshBlood   bool       `json:"freshBlood"`
	Inactive     bool       `json:"inactive"`
	Rank         string     `json:"rank"`
	SummonerID   string     `json:"summonerId"`
	LeaguePoints int        `json:"leaguePoints"`
}
