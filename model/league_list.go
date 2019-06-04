package model

import (
	"sort"
)

// LeagueList represents a league containing all player entries in it
type LeagueList struct {
	LeagueID      string        `json:"leagueId"`
	Tier          string        `json:"tier"`
	Entries       []LeagueEntry `json:"entries"`
	Queue         string        `json:"queue"`
	Name          string        `json:"name"`
	sortedEntries []LeagueEntry
}

// GetRank returns the entry at the given rank, sorted by league points
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
