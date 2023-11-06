package static

import (
	"fmt"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestClient_GetSeasons(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []Season
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]Season{{}}, 200),
			want: []Season{{}},
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
				c := NewClient(tt.doer, log.StandardLogger())
				got, err := c.GetSeasons()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetSeasons()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetQueues(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []Queue
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]Queue{{}}, 200),
			want: []Queue{{}},
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
				c := NewClient(tt.doer, log.StandardLogger())
				got, err := c.GetQueues()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetQueues()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetMaps(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []Map
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]Map{{}}, 200),
			want: []Map{{}},
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
				c := NewClient(tt.doer, log.StandardLogger())
				got, err := c.GetMaps()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetMaps()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetGameModes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []GameMode
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]GameMode{{}}, 200),
			want: []GameMode{{}},
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
				c := NewClient(tt.doer, log.StandardLogger())
				got, err := c.GetGameModes()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetGameModes()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetGameTypes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []GameType
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]GameType{{}}, 200),
			want: []GameType{{}},
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
				c := NewClient(tt.doer, log.StandardLogger())
				got, err := c.GetGameTypes()
				assert.Equal(t, tt.wantErr, err)
				if tt.wantErr == nil {
					assert.Equal(t, tt.want, got)
					got, err := c.GetGameTypes()
					assert.Nil(t, err)
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

func TestClient_GetGameMode(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      string
		want    GameMode
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]GameMode{{Mode: "id"}}, 200),
			id:   "id",
			want: GameMode{Mode: "id"},
		},
		{
			name:    "not found",
			doer:    mock.NewJSONMockDoer([]GameMode{}, 200),
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
				client := NewClient(test.doer, log.StandardLogger())
				got, err := client.GetGameMode(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetGameType(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      string
		want    GameType
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]GameType{{Type: "id"}}, 200),
			id:   "id",
			want: GameType{Type: "id"},
		},
		{
			name:    "not found",
			doer:    mock.NewJSONMockDoer([]GameType{}, 200),
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
				client := NewClient(test.doer, log.StandardLogger())
				got, err := client.GetGameType(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetMap(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      int
		want    Map
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]Map{{ID: 1}}, 200),
			id:   1,
			want: Map{ID: 1},
		},
		{
			name:    "not found",
			doer:    mock.NewJSONMockDoer([]Map{}, 200),
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
				client := NewClient(test.doer, log.StandardLogger())
				got, err := client.GetMap(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetQueue(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      int
		want    Queue
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]Queue{{ID: 1}}, 200),
			id:   1,
			want: Queue{ID: 1},
		},
		{
			name:    "not found",
			doer:    mock.NewJSONMockDoer([]Queue{}, 200),
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
				client := NewClient(test.doer, log.StandardLogger())
				got, err := client.GetQueue(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_GetSeason(t *testing.T) {
	type test struct {
		name    string
		doer    internal.Doer
		id      int
		want    Season
		wantErr error
	}
	tests := []test{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]Season{{ID: 1}}, 200),
			id:   1,
			want: Season{ID: 1},
		},
		{
			name:    "not found",
			doer:    mock.NewJSONMockDoer([]Season{}, 200),
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
				client := NewClient(test.doer, log.StandardLogger())
				got, err := client.GetSeason(test.id)
				assert.Equal(t, test.wantErr, err)
				assert.Equal(t, test.want, got)
			},
		)
	}
}

func TestClient_getInto(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		target  interface{}
		doer    internal.Doer
		wantErr bool
	}{
		{
			name: "fail do",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name:    "fail decode",
			target:  struct{}{},
			doer:    mock.NewJSONMockDoer(0, 200),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := NewClient(tt.doer, log.StandardLogger())
				err := c.getInto("endpoint", tt.target)
				assert.Equal(t, tt.wantErr, err != nil)
			},
		)
	}
}

func TestClient_ClearCaches(t *testing.T) {
	client := NewClient(http.DefaultClient, log.StandardLogger())
	client.ClearCaches()
}
