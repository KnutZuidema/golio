package riot

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLeagueList_GetRank(t *testing.T) {
	tests := []struct {
		name    string
		i       int
		entries []LeagueItem
		want    LeagueItem
	}{
		{
			name: "get rank",
			i:    0,
			entries: []LeagueItem{
				{
					LeaguePoints: 0,
				},
				{
					LeaguePoints: 1,
				},
			},
			want: LeagueItem{
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
