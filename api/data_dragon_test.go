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
			name:    "invalid data dragon response",
			doer:    mock.NewJSONMockDoer(struct{}{}, 200),
			wantErr: fmt.Errorf("response does not contain requested champion data"),
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
		format   dataDragonURL
		wantErr  bool
	}{
		{
			name:     "invalid url",
			endpoint: "a\nb c@asd*asd)",
			doer:     mock.NewStatusMockDoer(200),
			format:   dataDragonDataURLFormat,
			wantErr:  true,
		},
		{
			name: "error doer",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("error")
				},
			},
			format:  dataDragonImageURLFormat,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(tt.doer, RegionEuropeWest, log.StandardLogger())
			_, err := c.doRequest(tt.format, tt.endpoint)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestDataDragonClient_init(t *testing.T) {
	tests := []struct {
		name    string
		doer    Doer
		wantErr bool
	}{
		{
			name: "error on decode",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return &http.Response{
						Body: errorReadCloser{},
					}, nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(tt.doer, RegionOceania, log.StandardLogger())
			if err := c.init(RegionOceania); (err != nil) != tt.wantErr {
				t.Errorf("DataDragonClient.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataDragonClient_getInto(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "fail decode",
			target:  failJSONDecoding{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewDataDragonClient(mock.NewJSONMockDoer(0, 200), RegionOceania, log.StandardLogger())
			err := c.getInto("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_versionGreaterThan(t *testing.T) {
	type args struct {
		v1 string
		v2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid second arg",
			args: args{
				v1: "1",
				v2: "a",
			},
			want: false,
		},
		{
			name: "second greater",
			args: args{
				v1: "1",
				v2: "2",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := versionGreaterThan(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("versionGreaterThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func dataDragonResponseDoer(object interface{}) Doer {
	return mock.NewJSONMockDoer(model.DataDragonResponse{
		Data: object,
	}, 200)
}

type errorReadCloser struct{}

func (e errorReadCloser) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func (e errorReadCloser) Close() error {
	return fmt.Errorf("error")
}
