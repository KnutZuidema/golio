package lol

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigithanbalci/golio/api"
	"github.com/yigithanbalci/golio/internal"
	"github.com/yigithanbalci/golio/internal/mock"
	"net/http"
	"testing"
)

func TestChallengesClient_GetConfig(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*ChallengeConfigInfoDto
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*ChallengeConfigInfoDto{{}},
			doer: mock.NewJSONMockDoer([]ChallengeConfigInfoDto{{}}, 200),
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
			got, err := (&ChallengesClient{c: client}).GetConfig()
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestChallengesClient_GetPercentiles(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    map[string]map[string]float64
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: map[string]map[string]float64{},
			doer: mock.NewJSONMockDoer(map[string]map[string]float64{}, 200),
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
			got, err := (&ChallengesClient{c: client}).GetPercentiles()
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestChallengesClient_GetConfigWithChallengeId(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *ChallengeConfigInfoDto
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &ChallengeConfigInfoDto{},
			doer: mock.NewJSONMockDoer(ChallengeConfigInfoDto{}, 200),
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
			got, err := (&ChallengesClient{c: client}).GetConfigWithChallengeId(1)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestChallengesClient_GetLeaderBoardByChallengeIdAndLevel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*ApexPlayerInfoDto
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*ApexPlayerInfoDto{{}},
			doer: mock.NewJSONMockDoer([]ApexPlayerInfoDto{{}}, 200),
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
			got, err := (&ChallengesClient{c: client}).GetLeaderBoardByChallengeIdAndLevel(203102, TierMaster, 15)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestChallengesClient_GetPercentilesWithChallengeId(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    map[string]float64
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: map[string]float64{},
			doer: mock.NewJSONMockDoer(map[string]float64{}, 200),
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
			got, err := (&ChallengesClient{c: client}).GetPercentilesWithChallengeId(1)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestChallengesClient_GetPlayerDataWithPUUID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *PlayerInfoDto
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &PlayerInfoDto{},
			doer: mock.NewJSONMockDoer(PlayerInfoDto{}, 200),
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
			got, err := (&ChallengesClient{c: client}).GetPlayerDataWithPUUID("1")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
