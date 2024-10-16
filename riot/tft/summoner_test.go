package tft

import (
	"fmt"
	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestTFTSummoner_GetSummonerByAccountID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Summoner
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Summoner{},
			doer: mock.NewJSONMockDoer(Summoner{}, 200),
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
				got, err := (&SummonerClient{c: client}).GetSummonerByAccountID("accountId")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestTFTSummoner_GetSummonerByPUUID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Summoner
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Summoner{},
			doer: mock.NewJSONMockDoer(Summoner{}, 200),
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
				got, err := (&SummonerClient{c: client}).GetSummonerByPUUID("puuid")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestTFTSummoner_GetSummonerByMe(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Summoner
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Summoner{},
			doer: mock.NewJSONMockDoer(Summoner{}, 200),
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
				got, err := (&SummonerClient{c: client}).GetSummonerByMe("token")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestTFTSummoner_GetSummonerBySummonerID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Summoner
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Summoner{},
			doer: mock.NewJSONMockDoer(Summoner{}, 200),
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
				got, err := (&SummonerClient{c: client}).GetSummonerBySummonerID("summonerID")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}
