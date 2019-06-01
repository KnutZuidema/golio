package model

import (
	"sort"
)

type LeagueList struct {
	LeagueID      string        `json:"leagueId"`
	Tier          string        `json:"tier"`
	Entries       []LeagueEntry `json:"entries"`
	Queue         string        `json:"queue"`
	Name          string        `json:"name"`
	sortedEntries []LeagueEntry
}

func (l *LeagueList) GetRank(i int) LeagueEntry {
	if l.sortedEntries == nil || len(l.sortedEntries) != len(l.Entries) {
		l.sortedEntries = make([]LeagueEntry, len(l.Entries))
		copy(l.sortedEntries, l.Entries)
		sort.Slice(l.sortedEntries, func(i, j int) bool {
			return l.sortedEntries[i].LeaguePoints > l.sortedEntries[j].LeaguePoints
		})
	}
	return l.sortedEntries[i]
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
