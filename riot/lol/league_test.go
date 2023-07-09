package lol

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/yigithanbalci/golio/api"
	"github.com/yigithanbalci/golio/internal"
	"github.com/yigithanbalci/golio/internal/mock"
)

func TestLeagueClient_GetChallenger(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *LeagueList
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &LeagueList{},
			doer: mock.NewJSONMockDoer(LeagueList{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&LeagueClient{c: client}).GetChallenger(QueueRankedSolo)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestLeagueClient_GetGrandmaster(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *LeagueList
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &LeagueList{},
			doer: mock.NewJSONMockDoer(LeagueList{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&LeagueClient{c: client}).GetGrandmaster(QueueRankedSolo)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestLeagueClient_GetMaster(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *LeagueList
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &LeagueList{},
			doer: mock.NewJSONMockDoer(LeagueList{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&LeagueClient{c: client}).GetMaster(QueueRankedSolo)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestLeagueClient_ListPlayers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*LeagueItem
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*LeagueItem{},
			doer: mock.NewJSONMockDoer([]*LeagueItem{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&LeagueClient{c: client}).ListPlayers(QueueRankedSolo, TierGold, DivisionOne)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestLeagueClient_ListBySummoner(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*LeagueItem
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*LeagueItem{},
			doer: mock.NewJSONMockDoer([]*LeagueItem{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&LeagueClient{c: client}).ListBySummoner("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestLeagueClient_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *LeagueList
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &LeagueList{},
			doer: mock.NewJSONMockDoer(LeagueList{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&LeagueClient{c: client}).Get("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
