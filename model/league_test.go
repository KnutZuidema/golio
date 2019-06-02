package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLeagueList_GetRank(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		entries []LeagueEntry
		want    LeagueEntry
	}{
		{
			name: "get rank",
			i:    0,
			entries: []LeagueEntry{
				{
					LeaguePoints: 0,
				},
				{
					LeaguePoints: 1,
				},
			},
			want: LeagueEntry{
				LeaguePoints: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LeagueList{
				Entries: tt.entries,
			}
			require.Equal(t, tt.want, l.GetRank(tt.i))
		})
	}
}
