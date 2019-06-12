package api

import (
	"fmt"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/mock"
	"github.com/KnutZuidema/golio/model"
)

func TestNewDataDragonClient(t *testing.T) {
	ddClient := NewDataDragonClient(http.DefaultClient, RegionEuropeWest, log.StandardLogger())
	require.NotNil(t, ddClient)
}

func TestDataDragonClient_GetChampions(t *testing.T) {
	tests := []struct {
		name    string
		doer    Doer
		want    map[string]model.ChampionData
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.ChampionData{}),
			want: map[string]model.ChampionData{},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(tt.doer, RegionEuropeWest, log.StandardLogger())
			got, err := c.GetChampions()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetChampion(t *testing.T) {
	tests := []struct {
		name    string
		doer    Doer
		want    *model.ChampionDataExtended
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]*model.ChampionDataExtended{
				"champion": {},
			}),
			want: &model.ChampionDataExtended{},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(tt.doer, RegionEuropeWest, log.StandardLogger())
			got, err := c.GetChampion("champion")
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetProfileIcons(t *testing.T) {
	tests := []struct {
		name    string
		doer    Doer
		want    map[string]model.ProfileIcon
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.ProfileIcon{}),
			want: map[string]model.ProfileIcon{},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(tt.doer, RegionEuropeWest, log.StandardLogger())
			got, err := c.GetProfileIcons()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetItems(t *testing.T) {
	tests := []struct {
		name    string
		doer    Doer
		want    map[string]model.Item
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.Item{}),
			want: map[string]model.Item{},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(tt.doer, RegionEuropeWest, log.StandardLogger())
			got, err := c.GetItems()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetMasteries(t *testing.T) {
	tests := []struct {
		name    string
		doer    Doer
		want    map[string]model.Mastery
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.Mastery{}),
			want: map[string]model.Mastery{},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(tt.doer, RegionEuropeWest, log.StandardLogger())
			got, err := c.GetMasteries()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetSummonerSpells(t *testing.T) {
	tests := []struct {
		name    string
		doer    Doer
		want    map[string]model.SummonerSpell
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.SummonerSpell{}),
			want: map[string]model.SummonerSpell{},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(tt.doer, RegionEuropeWest, log.StandardLogger())
			got, err := c.GetSummonerSpells()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_doRequest(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		doer     Doer
		wantErr  bool
	}{
		{
			name:     "invalid url",
			endpoint: "a\nb c@asd*asd)",
			doer:     mock.NewStatusMockDoer(200),
			wantErr:  true,
		},
		{
			name: "error doer",
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
			c := NewDataDragonClient(tt.doer, RegionEuropeWest, log.StandardLogger())
			_, err := c.doRequest(dataDragonDataURLFormat, tt.endpoint)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func dataDragonResponseDoer(object interface{}) Doer {
	return mock.NewJSONMockDoer(model.DataDragonResponse{
		Data: object,
	}, 200)
}
