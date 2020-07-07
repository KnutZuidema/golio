package lol

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestChampionMasteryClient_List(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*ChampionMastery
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*ChampionMastery{},
			doer: mock.NewJSONMockDoer([]*ChampionMastery{}, 200),
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
			want: []*ChampionMastery{},
			doer: rateLimitDoer([]*ChampionMastery{}),
		},
		{
			name: "unavailable once",
			want: []*ChampionMastery{},
			doer: unavailableOnceDoer([]*ChampionMastery{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&ChampionMasteryClient{Client: client}).List("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestChampionMasteryClient_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *ChampionMastery
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &ChampionMastery{},
			doer: mock.NewJSONMockDoer(&ChampionMastery{}, 200),
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
			want: &ChampionMastery{},
			doer: rateLimitDoer(&ChampionMastery{}),
		},
		{
			name: "unavailable once",
			want: &ChampionMastery{},
			doer: unavailableOnceDoer(&ChampionMastery{}),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&ChampionMasteryClient{Client: client}).Get("id", "id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestChampionMasteryClient_GetTotal(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    int
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: 1,
			doer: mock.NewJSONMockDoer(1, 200),
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
			want: 1,
			doer: rateLimitDoer(1),
		},
		{
			name: "unavailable once",
			want: 1,
			doer: unavailableOnceDoer(1),
		},
		{
			name:    "unavailable twice",
			wantErr: api.ErrServiceUnavailable,
			doer:    mock.NewStatusMockDoer(http.StatusServiceUnavailable),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&ChampionMasteryClient{Client: client}).GetTotal("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func rateLimitDoer(object interface{}) internal.Doer {
	rateLimitCount := 0
	return &mock.Doer{
		Custom: func(r *http.Request) (*http.Response, error) {
			if rateLimitCount == 1 {
				return mock.NewJSONMockDoer(object, 200).Do(r)
			}
			rateLimitCount++
			return mock.NewHeaderMockDoer(http.StatusTooManyRequests, http.Header{
				"Retry-After": []string{"1"},
			}).Do(r)
		},
	}
}

func unavailableOnceDoer(object interface{}) internal.Doer {
	unavailableCount := 0
	return &mock.Doer{
		Custom: func(r *http.Request) (*http.Response, error) {
			if unavailableCount == 1 {
				return mock.NewJSONMockDoer(object, 200).Do(r)
			}
			unavailableCount++
			return mock.NewStatusMockDoer(http.StatusServiceUnavailable).Do(r)
		},
	}
}