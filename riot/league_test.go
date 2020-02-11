package riot

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
			name: "unknown error status",
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &LeagueList{},
			doer: rateLimitDoer(LeagueList{}),
		},
		{
			name: "unavailable once",
			want: &LeagueList{},
			doer: unavailableOnceDoer(LeagueList{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.League.GetChallenger(QueueRankedSolo)
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
			name: "unknown error status",
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &LeagueList{},
			doer: rateLimitDoer(LeagueList{}),
		},
		{
			name: "unavailable once",
			want: &LeagueList{},
			doer: unavailableOnceDoer(LeagueList{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.League.GetGrandmaster(QueueRankedSolo)
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
			name: "unknown error status",
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &LeagueList{},
			doer: rateLimitDoer(LeagueList{}),
		},
		{
			name: "unavailable once",
			want: &LeagueList{},
			doer: unavailableOnceDoer(LeagueList{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.League.GetMaster(QueueRankedSolo)
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
			name: "unknown error status",
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: []*LeagueItem{},
			doer: rateLimitDoer([]*LeagueItem{}),
		},
		{
			name: "unavailable once",
			want: []*LeagueItem{},
			doer: unavailableOnceDoer([]*LeagueItem{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.League.ListPlayers(QueueRankedSolo, TierGold, DivisionOne)
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
			name: "unknown error status",
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: []*LeagueItem{},
			doer: rateLimitDoer([]*LeagueItem{}),
		},
		{
			name: "unavailable once",
			want: []*LeagueItem{},
			doer: unavailableOnceDoer([]*LeagueItem{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.League.ListBySummoner("id")
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
			name: "unknown error status",
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &LeagueList{},
			doer: rateLimitDoer(LeagueList{}),
		},
		{
			name: "unavailable once",
			want: &LeagueList{},
			doer: unavailableOnceDoer(LeagueList{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.League.Get("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
