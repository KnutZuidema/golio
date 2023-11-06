package val

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestChallengesClient_GetMatchById(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Match
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Match{},
			doer: mock.NewJSONMockDoer(Match{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&MatchClient{c: client}).GetMatchByID("match-id")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestChallengesClient_GetMatchListByPUUID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *MatchList
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &MatchList{},
			doer: mock.NewJSONMockDoer(MatchList{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&MatchClient{c: client}).GetMatchListByPUUID("puuid")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestChallengesClient_GetRecentMatchesByQueue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *RecentMatches
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &RecentMatches{},
			doer: mock.NewJSONMockDoer(RecentMatches{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&MatchClient{c: client}).GetRecentMatchesByQueue("queue")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}
