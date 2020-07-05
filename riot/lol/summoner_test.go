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
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&SummonerClient{Client: client}).GetByName("name")
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
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&SummonerClient{Client: client}).GetByAccountID("accountID")
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
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&SummonerClient{Client: client}).GetByPUUID("puuid")
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
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&SummonerClient{Client: client}).GetByID("id")
			assert.Equal(t, err, tt.wantErr)
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
