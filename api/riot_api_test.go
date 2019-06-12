package api

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/mock"
	"github.com/KnutZuidema/golio/model"
)

func TestRiotAPIClient_GetSummonerByName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    Doer
		want    *model.Summoner
		wantErr error
	}{
		{
			name: "get repsonse",
			want: &model.Summoner{},
			doer: mock.NewJSONMockDoer(&model.Summoner{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.Summoner{},
			doer: rateLimitDoer(&model.Summoner{}),
		},
		{
			name: "unavailable once",
			want: &model.Summoner{},
			doer: unavailableOnceDoer(&model.Summoner{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		doer    Doer
		want    *model.Summoner
		wantErr error
	}{
		{
			name: "get response",
			want: &model.Summoner{},
			doer: mock.NewJSONMockDoer(&model.Summoner{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.Summoner{},
			doer: rateLimitDoer(&model.Summoner{}),
		},
		{
			name: "unavailable once",
			want: &model.Summoner{},
			doer: unavailableOnceDoer(&model.Summoner{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		doer    Doer
		want    *model.Summoner
		wantErr error
	}{
		{
			name: "get response",
			want: &model.Summoner{},
			doer: mock.NewJSONMockDoer(&model.Summoner{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.Summoner{},
			doer: rateLimitDoer(&model.Summoner{}),
		},
		{
			name: "unavailable once",
			want: &model.Summoner{},
			doer: unavailableOnceDoer(&model.Summoner{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		doer    Doer
		want    *model.Summoner
		wantErr error
	}{
		{
			name: "get response",
			want: &model.Summoner{},
			doer: mock.NewJSONMockDoer(&model.Summoner{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.Summoner{},
			doer: rateLimitDoer(&model.Summoner{}),
		},
		{
			name: "unavailable once",
			want: &model.Summoner{},
			doer: unavailableOnceDoer(&model.Summoner{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    []*model.ChampionMastery
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*model.ChampionMastery{},
			doer: mock.NewJSONMockDoer([]*model.ChampionMastery{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: []*model.ChampionMastery{},
			doer: rateLimitDoer([]*model.ChampionMastery{}),
		},
		{
			name: "unavailable once",
			want: []*model.ChampionMastery{},
			doer: unavailableOnceDoer([]*model.ChampionMastery{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.ChampionMastery
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.ChampionMastery{},
			doer: mock.NewJSONMockDoer(&model.ChampionMastery{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.ChampionMastery{},
			doer: rateLimitDoer(&model.ChampionMastery{}),
		},
		{
			name: "unavailable once",
			want: &model.ChampionMastery{},
			doer: unavailableOnceDoer(&model.ChampionMastery{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: 1,
			doer: mock.NewJSONMockDoer(1, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
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
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.ChampionInfo
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.ChampionInfo{},
			doer: mock.NewJSONMockDoer(model.ChampionInfo{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.ChampionInfo{},
			doer: rateLimitDoer(model.ChampionInfo{}),
		},
		{
			name: "unavailable once",
			want: &model.ChampionInfo{},
			doer: unavailableOnceDoer(model.ChampionInfo{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.LeagueList
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.LeagueList{},
			doer: mock.NewJSONMockDoer(model.LeagueList{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.LeagueList{},
			doer: rateLimitDoer(model.LeagueList{}),
		},
		{
			name: "unavailable once",
			want: &model.LeagueList{},
			doer: unavailableOnceDoer(model.LeagueList{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.LeagueList
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.LeagueList{},
			doer: mock.NewJSONMockDoer(model.LeagueList{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.LeagueList{},
			doer: rateLimitDoer(model.LeagueList{}),
		},
		{
			name: "unavailable once",
			want: &model.LeagueList{},
			doer: unavailableOnceDoer(model.LeagueList{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.LeagueList
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.LeagueList{},
			doer: mock.NewJSONMockDoer(model.LeagueList{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.LeagueList{},
			doer: rateLimitDoer(model.LeagueList{}),
		},
		{
			name: "unavailable once",
			want: &model.LeagueList{},
			doer: unavailableOnceDoer(model.LeagueList{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.GetMasterLeague(QueueRankedSolo)
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
		want    []*model.LeagueEntry
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*model.LeagueEntry{},
			doer: mock.NewJSONMockDoer([]*model.LeagueEntry{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: []*model.LeagueEntry{},
			doer: rateLimitDoer([]*model.LeagueEntry{}),
		},
		{
			name: "unavailable once",
			want: []*model.LeagueEntry{},
			doer: unavailableOnceDoer([]*model.LeagueEntry{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    []*model.LeagueEntry
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*model.LeagueEntry{},
			doer: mock.NewJSONMockDoer([]*model.LeagueEntry{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: []*model.LeagueEntry{},
			doer: rateLimitDoer([]*model.LeagueEntry{}),
		},
		{
			name: "unavailable once",
			want: []*model.LeagueEntry{},
			doer: unavailableOnceDoer([]*model.LeagueEntry{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.LeagueList
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.LeagueList{},
			doer: mock.NewJSONMockDoer(model.LeagueList{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.LeagueList{},
			doer: rateLimitDoer(model.LeagueList{}),
		},
		{
			name: "unavailable once",
			want: &model.LeagueList{},
			doer: unavailableOnceDoer(model.LeagueList{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.Status
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.Status{},
			doer: mock.NewJSONMockDoer(model.Status{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.Status{},
			doer: rateLimitDoer(model.Status{}),
		},
		{
			name: "unavailable once",
			want: &model.Status{},
			doer: unavailableOnceDoer(model.Status{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.Matchlist
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.Matchlist{},
			doer: mock.NewJSONMockDoer(model.Matchlist{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.Matchlist{},
			doer: rateLimitDoer(model.Matchlist{}),
		},
		{
			name: "unavailable once",
			want: &model.Matchlist{},
			doer: unavailableOnceDoer(model.Matchlist{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
	tests := []struct {
		name    string
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer(model.Matchlist{
				Matches: []model.MatchReference{
					{},
				},
			}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			doer: rateLimitDoer(model.Matchlist{}),
		},
		{
			name: "unavailable once",
			doer: unavailableOnceDoer(model.Matchlist{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got := client.GetMatchesByAccountStream("id")
			for res := range got {
				if res.Error != nil && tt.wantErr != nil {
					require.Equal(t, res.Error, tt.wantErr)
					break
				} else if res.Error != nil {
					require.Equal(t, res.Error, io.EOF)
					return
				}
				assert.NotNil(t, res.MatchReference)
			}
		})
	}
}

func TestRiotAPIClient_GetMatch(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *model.Match
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.Match{},
			doer: mock.NewJSONMockDoer(model.Match{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.Match{},
			doer: rateLimitDoer(model.Match{}),
		},
		{
			name: "unavailable once",
			want: &model.Match{},
			doer: unavailableOnceDoer(model.Match{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.GetMatch(1)
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
		want    *model.FeaturedGames
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.FeaturedGames{},
			doer: mock.NewJSONMockDoer(model.FeaturedGames{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.FeaturedGames{},
			doer: rateLimitDoer(model.FeaturedGames{}),
		},
		{
			name: "unavailable once",
			want: &model.FeaturedGames{},
			doer: unavailableOnceDoer(model.FeaturedGames{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
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
		want    *model.CurrentGameInfo
		doer    Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &model.CurrentGameInfo{},
			doer: mock.NewJSONMockDoer(model.CurrentGameInfo{}, 200),
		},
		{
			name: "unknown error status",
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
			doer: mock.NewStatusMockDoer(999),
		},
		{
			name:    "not found",
			wantErr: ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
		{
			name: "rate limited",
			want: &model.CurrentGameInfo{},
			doer: rateLimitDoer(model.CurrentGameInfo{}),
		},
		{
			name: "unavailable once",
			want: &model.CurrentGameInfo{},
			doer: unavailableOnceDoer(model.CurrentGameInfo{}),
		},
		{
			name:    "unavailable twice",
			wantErr: ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewRiotAPIClient(RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := client.GetCurrentGame("id")
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
		body     string
	}
	tests := []struct {
		name    string
		args    args
		doer    Doer
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewRiotAPIClient(RegionEuropeNorthEast, "", tt.doer, logrus.StandardLogger())
			_, err := c.doRequest(tt.args.method, tt.args.endpoint, tt.args.body)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func rateLimitDoer(object interface{}) Doer {
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

func unavailableOnceDoer(object interface{}) Doer {
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
