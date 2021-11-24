package lol

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestMatchClient_List(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []string
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []string{},
			doer: mock.NewJSONMockDoer([]string{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			queue := 200
			got, err := (&MatchClient{c: client}).List("id", 0, 1, &MatchListOptions{
				Queue:     &queue,
				Type:      "SomeType",
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour),
			})
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestMatchClient_ListStream(t *testing.T) {
	t.Parallel()
	count := 0
	tests := []struct {
		name    string
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					if count == 0 {
						count++
						return mock.NewJSONMockDoer(make([]string, 100), 200).Do(r)
					}
					return mock.NewJSONMockDoer([]string{}, 200).Do(r)
				},
			},
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			queue := 200
			got := (&MatchClient{c: client}).ListStream("id", &MatchListOptions{
				Queue:     &queue,
				Type:      "SomeType",
				StartTime: time.Now(),
				EndTime:   time.Now().Add(time.Hour),
			})
			for res := range got {
				if res.Error != nil && tt.wantErr != nil {
					require.Equal(t, res.Error, tt.wantErr)
					break
				} else if res.Error != nil {
					require.Equal(t, res.Error, io.EOF)
					return
				}
			}
		})
	}
}

func TestMatchClient_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Match
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Match{},
			doer: mock.NewJSONMockDoer(Match{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&MatchClient{c: client}).Get("NA_1")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestMatchClient_GetTimeline(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *MatchTimeline
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &MatchTimeline{},
			doer: mock.NewJSONMockDoer(MatchTimeline{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&MatchClient{c: client}).GetTimeline(0)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
