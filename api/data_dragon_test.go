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
	t.Parallel()
	ddClient := NewDataDragonClient(http.DefaultClient, RegionEuropeWest, log.StandardLogger())
	require.NotNil(t, ddClient)
}

func TestDataDragonClient_GetChampions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    Doer
		want    []model.ChampionData
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.ChampionData{
				"champion": {},
			}),
			want: []model.ChampionData{{}},
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
				got, err := c.GetChampions()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetChampion(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    Doer
		want    model.ChampionDataExtended
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.ChampionDataExtended{
				"champion": {},
			}),
			want: model.ChampionDataExtended{},
		},
		{
			name:    "invalid data dragon response",
			doer:    mock.NewJSONMockDoer(struct{}{}, 200),
			wantErr: fmt.Errorf("no data for champion champion"),
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
				got, err := c.GetChampion("champion")
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetProfileIcons(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    Doer
		want    []model.ProfileIcon
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.ProfileIcon{
				"icon": {},
			}),
			want: []model.ProfileIcon{{}},
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
				got, err := c.GetProfileIcons()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetItems(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    Doer
		want    []model.Item
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.Item{
				"item": {},
			}),
			want: []model.Item{{ID: "item"}},
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
				got, err := c.GetItems()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetRunes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    Doer
		want    []model.Item
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.Item{
				"rune": {},
			}),
			want: []model.Item{{ID: "rune"}},
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
			got, err := c.GetRunes()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
				got, err := c.GetRunes()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetMasteries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    Doer
		want    []model.Mastery
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.Mastery{
				"mastery": {},
			}),
			want: []model.Mastery{{}},
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
				got, err := c.GetMasteries()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_GetSummonerSpells(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    Doer
		want    []model.SummonerSpell
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(map[string]model.SummonerSpell{
				"summoner": {},
			}),
			want: []model.SummonerSpell{{}},
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
				got, err := c.GetSummonerSpells()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDataDragonClient_ClearCaches(t *testing.T) {
	t.Parallel()
	c := NewDataDragonClient(http.DefaultClient, RegionKorea, log.StandardLogger())
	c.ClearCaches()
}

func TestDataDragonClient_doRequest(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
	t.Parallel()
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
