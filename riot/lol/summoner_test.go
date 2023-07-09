package lol

import (
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestSummonerClient_GetByName(t *testing.T) {
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
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&SummonerClient{c: client}).GetByName("name")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestSummonerClient_GetByAccountID(t *testing.T) {
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
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&SummonerClient{c: client}).GetByAccountID("accountID")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestSummonerClient_GetByPUUID(t *testing.T) {
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
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&SummonerClient{c: client}).GetByPUUID("puuid")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestSummonerClient_GetByID(t *testing.T) {
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
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&SummonerClient{c: client}).GetByID("id")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
