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
	"github.com/KnutZuidema/golio/model"
)

func TestStaticDataClient_GetSeasons(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []model.Season
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]model.Season{{}}, 200),
			want: []model.Season{{}},
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
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.doer, log.StandardLogger())
			got, err := c.GetSeasons()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
				got, err := c.GetSeasons()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStaticDataClient_GetQueues(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []model.Queue
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]model.Queue{{}}, 200),
			want: []model.Queue{{}},
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
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.doer, log.StandardLogger())
			got, err := c.GetQueues()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
				got, err := c.GetQueues()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStaticDataClient_GetMaps(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []model.Map
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]model.Map{{}}, 200),
			want: []model.Map{{}},
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
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.doer, log.StandardLogger())
			got, err := c.GetMaps()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
				got, err := c.GetMaps()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStaticDataClient_GetGameModes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []model.GameMode
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]model.GameMode{{}}, 200),
			want: []model.GameMode{{}},
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
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.doer, log.StandardLogger())
			got, err := c.GetGameModes()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
				got, err := c.GetGameModes()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStaticDataClient_GetGameTypes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		want    []model.GameType
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewJSONMockDoer([]model.GameType{{}}, 200),
			want: []model.GameType{{}},
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
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.doer, log.StandardLogger())
			got, err := c.GetGameTypes()
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
				got, err := c.GetGameTypes()
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStaticDataClient_getInto(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.doer, log.StandardLogger())
			err := c.getInto("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestStaticDataClient_ClearCaches(t *testing.T) {
	client := NewClient(http.DefaultClient, log.StandardLogger())
	client.ClearCaches()
}
