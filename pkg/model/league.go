package model

type LeagueList struct {
	LeagueID string        `json:"leagueId"`
	Tier     string        `json:"tier"`
	Entries  []LeagueEntry `json:"entries"`
	Queue    string        `json:"queue"`
	Name     string        `json:"name"`
}

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

type MiniSeries struct {
	Progress string `json:"progress"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}
