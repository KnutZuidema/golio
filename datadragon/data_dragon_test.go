package datadragon

import (
	"fmt"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestNewClient(t *testing.T) {
	t.Parallel()
	ddClient := NewClient(http.DefaultClient, api.RegionEuropeWest, log.StandardLogger())
	require.NotNil(t, ddClient)
}

func TestClient_GetChampions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []ChampionData
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]ChampionData{
					"champion": {},
				},
			),
			want: []ChampionData{{}},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: api.ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := c.GetChampions()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetChampions()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetChampion(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    ChampionDataExtended
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewPathJSONMockDoer(
				[]mock.PathJSONResponse{
					{
						PathSuffix: "/champion.json",
						Object: dataDragonResponse{
							Data: map[string]ChampionData{
								"champion-id": {
									ID:   "champion-id",
									Name: "champion-name",
								},
							},
						},
						Code: 200,
					},
					{
						PathSuffix: "/champion/champion-id.json",
						Object: dataDragonResponse{
							Data: map[string]ChampionDataExtended{
								"champion-id": {
									ChampionData: ChampionData{
										ID:   "champion-id",
										Name: "champion-name",
									},
								},
							},
						},
						Code: 200,
					},
				},
			),
			want: ChampionDataExtended{
				ChampionData: ChampionData{
					ID:   "champion-id",
					Name: "champion-name",
				},
			},
		},
		{
			name:    "not found",
			doer:    mock.NewJSONMockDoer(struct{}{}, 200),
			wantErr: api.ErrNotFound,
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: api.ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := c.GetChampion("champion-name")
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetChampion("champion-name")
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetProfileIcons(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []ProfileIcon
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]ProfileIcon{
					"icon": {},
				},
			),
			want: []ProfileIcon{{}},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: api.ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := c.GetProfileIcons()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetProfileIcons()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetItems(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []Item
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]Item{
					"item": {},
				},
			),
			want: []Item{{ID: "item"}},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: api.ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := c.GetItems()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetItems()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetRunes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []Item
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]Item{
					"rune": {},
				},
			),
			want: []Item{{ID: "rune"}},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: api.ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := c.GetRunes()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetRunes()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetMasteries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []Mastery
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]Mastery{
					"mastery": {},
				},
			),
			want: []Mastery{{}},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: api.ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := c.GetMasteries()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetMasteries()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetSummonerSpells(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []SummonerSpell
		wantErr error
	}{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]SummonerSpell{
					"summoner": {},
				},
			),
			want: []SummonerSpell{{}},
		},
		{
			name:    "known error",
			doer:    mock.NewStatusMockDoer(http.StatusForbidden),
			wantErr: api.ErrForbidden,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := c.GetSummonerSpells()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetSummonerSpells()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_ClearCaches(t *testing.T) {
	t.Parallel()
	c := NewClient(http.DefaultClient, api.RegionKorea, log.StandardLogger())
	c.ClearCaches()
}

func TestClient_GetChampionByID(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      string
		want    ChampionDataExtended
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]ChampionData{
					"champion-id": {ID: "champion-id", Name: "champion-name", Key: "champion-id"},
				},
			),
			id: "champion-id",
			want: ChampionDataExtended{
				ChampionData: ChampionData{
					ID:   "champion-id",
					Name: "champion-name",
					Key:  "champion-id",
				},
			},
		},
		{
			name:    "not found",
			doer:    dataDragonResponseDoer(map[string]ChampionData{}),
			wantErr: api.ErrNotFound,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := NewClient(test.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := client.GetChampionByID(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetProfileIcon(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      int
		want    ProfileIcon
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]ProfileIcon{
					"icon": {ID: 1},
				},
			),
			id:   1,
			want: ProfileIcon{ID: 1},
		},
		{
			name:    "not found",
			doer:    dataDragonResponseDoer(map[string]ProfileIcon{}),
			wantErr: api.ErrNotFound,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := NewClient(test.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := client.GetProfileIcon(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetItem(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      string
		want    Item
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]Item{
					"id": {},
				},
			),
			id:   "id",
			want: Item{ID: "id"},
		},
		{
			name:    "not found",
			doer:    dataDragonResponseDoer(map[string]Item{}),
			wantErr: api.ErrNotFound,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := NewClient(test.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := client.GetItem(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetMastery(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      int
		want    Mastery
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]Mastery{
					"icon": {ID: 1},
				},
			),
			id:   1,
			want: Mastery{ID: 1},
		},
		{
			name:    "not found",
			doer:    dataDragonResponseDoer(map[string]Mastery{}),
			wantErr: api.ErrNotFound,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := NewClient(test.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := client.GetMastery(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetRune(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      string
		want    Item
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]Item{
					"id": {},
				},
			),
			id:   "id",
			want: Item{ID: "id"},
		},
		{
			name:    "not found",
			doer:    dataDragonResponseDoer(map[string]Item{}),
			wantErr: api.ErrNotFound,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := NewClient(test.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := client.GetRune(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetSummonerSpell(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      string
		want    SummonerSpell
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: dataDragonResponseDoer(
				map[string]SummonerSpell{
					"id": {Key: "id"},
				},
			),
			id:   "id",
			want: SummonerSpell{Key: "id"},
		},
		{
			name:    "not found",
			doer:    dataDragonResponseDoer(map[string]SummonerSpell{}),
			wantErr: api.ErrNotFound,
		},
		{
			name: "unknown error",
			doer: mock.NewStatusMockDoer(999),
			wantErr: api.Error{
				Message:    "unknown error reason",
				StatusCode: 999,
			},
		},
	}
	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				client := NewClient(test.doer, api.RegionEuropeWest, log.StandardLogger())
				got, err := client.GetSummonerSpell(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_doRequest(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		endpoint string
		doer     internal.Doer
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
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionEuropeWest, log.StandardLogger())
				_, err := c.doRequest(tt.format, tt.endpoint)
				assert.Equal(t, err != nil, tt.wantErr)
			},
		)
	}
}

func TestClient_init(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
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
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, api.RegionOceania, log.StandardLogger())
				if err := c.init(string(api.RegionOceania)); (err != nil) != tt.wantErr {
					t.Errorf("Client.init() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestClient_getInto(t *testing.T) {
	t.Parallel()
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
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(mock.NewJSONMockDoer(0, 200), api.RegionOceania, log.StandardLogger())
				err := c.getInto("endpoint", tt.target)
				assert.Equal(t, tt.wantErr, err != nil)
			},
		)
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
		t.Run(
			tt.name, func(t *testing.T) {
				if got := versionGreaterThan(tt.args.v1, tt.args.v2); got != tt.want {
					t.Errorf("versionGreaterThan() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func dataDragonResponseDoer(object interface{}) internal.Doer {
	return mock.NewJSONMockDoer(
		dataDragonResponse{
			Data: object,
		}, 200,
	)
}

type errorReadCloser struct{}

func (e errorReadCloser) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("error")
}

func (e errorReadCloser) Close() error {
	return fmt.Errorf("error")
}
