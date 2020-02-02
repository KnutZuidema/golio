package riot

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestRiotAPIClient_GetSummonerByName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    *Summoner
		wantErr error
	}{
		{
			name: "get response",
			want: &Summoner{},
			doer: mock.NewJSONMockDoer(&Summoner{}, 200),
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
			want: &Summoner{},
			doer: rateLimitDoer(&Summoner{}),
		},
		{
			name: "unavailable once",
			want: &Summoner{},
			doer: unavailableOnceDoer(&Summoner{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.GetSummonerByName("name")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetSummonerByAccount(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    *Summoner
		wantErr error
	}{
		{
			name: "get response",
			want: &Summoner{},
			doer: mock.NewJSONMockDoer(&Summoner{}, 200),
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
			want: &Summoner{},
			doer: rateLimitDoer(&Summoner{}),
		},
		{
			name: "unavailable once",
			want: &Summoner{},
			doer: unavailableOnceDoer(&Summoner{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.GetSummonerByAccount("accountID")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetSummonerByPUUID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    *Summoner
		wantErr error
	}{
		{
			name: "get response",
			want: &Summoner{},
			doer: mock.NewJSONMockDoer(&Summoner{}, 200),
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
			want: &Summoner{},
			doer: rateLimitDoer(&Summoner{}),
		},
		{
			name: "unavailable once",
			want: &Summoner{},
			doer: unavailableOnceDoer(&Summoner{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.GetSummonerByPUUID("puuid")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetSummonerBySummonerID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    *Summoner
		wantErr error
	}{
		{
			name: "get response",
			want: &Summoner{},
			doer: mock.NewJSONMockDoer(&Summoner{}, 200),
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
			want: &Summoner{},
			doer: rateLimitDoer(&Summoner{}),
		},
		{
			name: "unavailable once",
			want: &Summoner{},
			doer: unavailableOnceDoer(&Summoner{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.GetSummonerBySummonerID("id")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetChampionMasteries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*ChampionMastery
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*ChampionMastery{},
			doer: mock.NewJSONMockDoer([]*ChampionMastery{}, 200),
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
			want: []*ChampionMastery{},
			doer: rateLimitDoer([]*ChampionMastery{}),
		},
		{
			name: "unavailable once",
			want: []*ChampionMastery{},
			doer: unavailableOnceDoer([]*ChampionMastery{}),
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
			got, err := client.GetChampionMasteries("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetChampionMastery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *ChampionMastery
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &ChampionMastery{},
			doer: mock.NewJSONMockDoer(&ChampionMastery{}, 200),
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
			want: &ChampionMastery{},
			doer: rateLimitDoer(&ChampionMastery{}),
		},
		{
			name: "unavailable once",
			want: &ChampionMastery{},
			doer: unavailableOnceDoer(&ChampionMastery{}),
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
			got, err := client.GetChampionMastery("id", "id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetChampionMasteryTotalScore(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    int
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: 1,
			doer: mock.NewJSONMockDoer(1, 200),
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
			want: 1,
			doer: rateLimitDoer(1),
		},
		{
			name: "unavailable once",
			want: 1,
			doer: unavailableOnceDoer(1),
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
			got, err := client.GetChampionMasteryTotalScore("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetFreeChampionRotation(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *ChampionInfo
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &ChampionInfo{},
			doer: mock.NewJSONMockDoer(ChampionInfo{}, 200),
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
			want: &ChampionInfo{},
			doer: rateLimitDoer(ChampionInfo{}),
		},
		{
			name: "unavailable once",
			want: &ChampionInfo{},
			doer: unavailableOnceDoer(ChampionInfo{}),
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
			got, err := client.GetFreeChampionRotation()
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetChallengerLeague(t *testing.T) {
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
			got, err := client.GetChallengerLeague(QueueRankedSolo)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetGrandmasterLeague(t *testing.T) {
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
			got, err := client.GetGrandmasterLeague(QueueRankedSolo)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetMasterLeague(t *testing.T) {
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
			got, err := client.GetMasterLeague(QueueRankedSolo)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetLeaguesBySummoner(t *testing.T) {
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
			got, err := client.GetLeaguesBySummoner("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetSummonerLeagues(t *testing.T) {
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
			got, err := client.GetLeaguesBySummoner("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetLeagues(t *testing.T) {
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
			got, err := client.GetLeagues(QueueRankedSolo, TierGold, DivisionThree)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetLeague(t *testing.T) {
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
			got, err := client.GetLeague("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Status
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Status{},
			doer: mock.NewJSONMockDoer(Status{}, 200),
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
			want: &Status{},
			doer: rateLimitDoer(Status{}),
		},
		{
			name: "unavailable once",
			want: &Status{},
			doer: unavailableOnceDoer(Status{}),
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
			got, err := client.GetStatus()
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetMatchesByAccount(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Matchlist
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Matchlist{},
			doer: mock.NewJSONMockDoer(Matchlist{}, 200),
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
			want: &Matchlist{},
			doer: rateLimitDoer(Matchlist{}),
		},
		{
			name: "unavailable once",
			want: &Matchlist{},
			doer: unavailableOnceDoer(Matchlist{}),
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
			got, err := client.GetMatchesByAccount("id", 0, 1)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetMatchesByAccountStream(t *testing.T) {
	t.Parallel()
	count := 0
	tests := []struct {
		name    string
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					if count == 0 {
						count++
						return mock.NewJSONMockDoer(Matchlist{
							Matches: make([]*MatchReference, 100),
						}, 200).Do(r)
					}
					return mock.NewJSONMockDoer(Matchlist{}, 200).Do(r)
				},
			},
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
			doer: rateLimitDoer(Matchlist{}),
		},
		{
			name: "unavailable once",
			doer: unavailableOnceDoer(Matchlist{}),
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
			got := client.GetMatchesByAccountStream("id")
			for res := range got {
				if res.Error != nil && tt.wantErr != nil {
					require.Equal(t, res.Error, tt.wantErr)
					break
				} else if res.Error != nil {
					require.Equal(t, res.Error, io.EOF)
					return
				}
			}
		})
	}
}

func TestRiotAPIClient_GetMatch(t *testing.T) {
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
			want: &Match{},
			doer: rateLimitDoer(Match{}),
		},
		{
			name: "unavailable once",
			want: &Match{},
			doer: unavailableOnceDoer(Match{}),
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
			got, err := client.GetMatch(1)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetMatchTimeline(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *MatchTimeline
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &MatchTimeline{},
			doer: mock.NewJSONMockDoer(MatchTimeline{}, 200),
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
			want: &MatchTimeline{},
			doer: rateLimitDoer(MatchTimeline{}),
		},
		{
			name: "unavailable once",
			want: &MatchTimeline{},
			doer: unavailableOnceDoer(MatchTimeline{}),
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
			got, err := client.GetMatchTimeline(0)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetMatchIDsByTournamentCode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []int
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []int{},
			doer: mock.NewJSONMockDoer([]int{}, 200),
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
			want: []int{},
			doer: rateLimitDoer([]int{}),
		},
		{
			name: "unavailable once",
			want: []int{},
			doer: unavailableOnceDoer([]int{}),
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
			got, err := client.GetMatchIDsByTournamentCode("tournamentCode")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetMatchForTournament(t *testing.T) {
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
			want: &Match{},
			doer: rateLimitDoer(Match{}),
		},
		{
			name: "unavailable once",
			want: &Match{},
			doer: unavailableOnceDoer(Match{}),
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
			got, err := client.GetMatchForTournament(0, "tournamentCode")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetFeaturedGames(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *FeaturedGames
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &FeaturedGames{},
			doer: mock.NewJSONMockDoer(FeaturedGames{}, 200),
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
			want: &FeaturedGames{},
			doer: rateLimitDoer(FeaturedGames{}),
		},
		{
			name: "unavailable once",
			want: &FeaturedGames{},
			doer: unavailableOnceDoer(FeaturedGames{}),
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
			got, err := client.GetFeaturedGames()
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetCurrentGame(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *GameInfo
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &GameInfo{},
			doer: mock.NewJSONMockDoer(GameInfo{}, 200),
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
			want: &GameInfo{},
			doer: rateLimitDoer(GameInfo{}),
		},
		{
			name: "unavailable once",
			want: &GameInfo{},
			doer: unavailableOnceDoer(GameInfo{}),
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
			got, err := client.GetCurrentGame("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_CreateTournamentCodes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []string
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []string{},
			doer: mock.NewJSONMockDoer([]string{}, 200),
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
			want: []string{},
			doer: rateLimitDoer([]string{}),
		},
		{
			name: "unavailable once",
			want: []string{},
			doer: unavailableOnceDoer([]string{}),
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
			got, err := client.CreateTournamentCodes(0, 0, &TournamentCodeParameters{}, true)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetLobbyEvents(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *LobbyEventList
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &LobbyEventList{},
			doer: mock.NewJSONMockDoer(LobbyEventList{}, 200),
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
			want: &LobbyEventList{},
			doer: rateLimitDoer(LobbyEventList{}),
		},
		{
			name: "unavailable once",
			want: &LobbyEventList{},
			doer: unavailableOnceDoer(LobbyEventList{}),
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
			got, err := client.GetLobbyEvents("code", true)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_CreateTournamentProvider(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    int
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: 1,
			doer: mock.NewJSONMockDoer(1, 200),
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
			want: 1,
			doer: rateLimitDoer(1),
		},
		{
			name: "unavailable once",
			want: 1,
			doer: unavailableOnceDoer(1),
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
			got, err := client.CreateTournamentProvider(&ProviderRegistrationParameters{}, true)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_CreateTournament(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    int
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: 1,
			doer: mock.NewJSONMockDoer(1, 200),
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
			want: 1,
			doer: rateLimitDoer(1),
		},
		{
			name: "unavailable once",
			want: 1,
			doer: unavailableOnceDoer(1),
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
			got, err := client.CreateTournament(&TournamentRegistrationParameters{}, true)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_GetTournament(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Tournament
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Tournament{},
			doer: mock.NewJSONMockDoer(Tournament{}, 200),
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
			want: &Tournament{},
			doer: rateLimitDoer(Tournament{}),
		},
		{
			name: "unavailable once",
			want: &Tournament{},
			doer: unavailableOnceDoer(Tournament{}),
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
			got, err := client.GetTournament("code")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_UpdateTournament(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewStatusMockDoer(200),
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
			doer: rateLimitDoer(1),
		},
		{
			name: "unavailable once",
			doer: unavailableOnceDoer(1),
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
			err := client.UpdateTournament("code", TournamentUpdateParameters{})
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
		})
	}
}

func TestRiotAPIClient_GetThirdPartyCode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    string
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: "code",
			doer: mock.NewJSONMockDoer("code", 200),
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
			want: "code",
			doer: rateLimitDoer("code"),
		},
		{
			name: "unavailable once",
			want: "code",
			doer: unavailableOnceDoer("code"),
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
			got, err := client.GetThirdPartyCode("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestRiotAPIClient_doRequest(t *testing.T) {
	t.Parallel()
	type args struct {
		method   string
		endpoint string
		body     io.Reader
	}
	tests := []struct {
		name    string
		args    args
		doer    internal.Doer
		wantErr bool
	}{
		{
			name: "invalid method",
			args: args{
				method: "a\nb c",
			},
			doer:    mock.NewStatusMockDoer(200),
			wantErr: true,
		},
		{
			name: "error doer",
			args: args{
				method: "GET",
			},
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name: "fail second try",
			args: args{
				method: "GET",
			},
			doer:    failOnSecondDoer(),
			wantErr: true,
		},
		{
			name: "invalid retry header",
			args: args{
				method: "GET",
			},
			doer:    invalidHeaderDoer(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionEuropeNorthEast, "", tt.doer, logrus.StandardLogger())
			_, err := c.doRequest(tt.args.method, tt.args.endpoint, tt.args.body)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestRiotAPIClient_getInto(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "fail decode",
			target:  struct{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", mock.NewJSONMockDoer(0, 200), logrus.StandardLogger())
			err := c.getInto("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestRiotAPIClient_postInto(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "fail decode",
			target:  struct{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", mock.NewJSONMockDoer(0, 200), logrus.StandardLogger())
			err := c.postInto("endpoint", struct{}{}, tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestRiotAPIClient_post(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "fail encode",
			target:  mock.FailJSONEncoding{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", mock.NewStatusMockDoer(200), logrus.StandardLogger())
			_, err := c.post("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestRiotAPIClient_put(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "fail encode",
			target:  mock.FailJSONEncoding{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", mock.NewStatusMockDoer(200), logrus.StandardLogger())
			err := c.put("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func rateLimitDoer(object interface{}) internal.Doer {
	rateLimitCount := 0
	return &mock.Doer{
		Custom: func(r *http.Request) (*http.Response, error) {
			if rateLimitCount == 1 {
				return mock.NewJSONMockDoer(object, 200).Do(r)
			}
			rateLimitCount++
			return mock.NewHeaderMockDoer(http.StatusTooManyRequests, http.Header{
				"Retry-After": []string{"1"},
			}).Do(r)
		},
	}
}

func unavailableOnceDoer(object interface{}) internal.Doer {
	unavailableCount := 0
	return &mock.Doer{
		Custom: func(r *http.Request) (*http.Response, error) {
			if unavailableCount == 1 {
				return mock.NewJSONMockDoer(object, 200).Do(r)
			}
			unavailableCount++
			return mock.NewStatusMockDoer(http.StatusServiceUnavailable).Do(r)
		},
	}
}

func failOnSecondDoer() internal.Doer {
	count := 0
	return &mock.Doer{
		Custom: func(r *http.Request) (*http.Response, error) {
			if count == 0 {
				count++
				return mock.NewStatusMockDoer(http.StatusServiceUnavailable).Do(r)
			}
			return nil, fmt.Errorf("error")
		},
	}
}

func invalidHeaderDoer() internal.Doer {
	return mock.NewHeaderMockDoer(http.StatusTooManyRequests, http.Header{
		"Retry-After": []string{"abc"},
	})
}
