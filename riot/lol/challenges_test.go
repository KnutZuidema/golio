package lol

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigithanbalci/golio/api"
	"github.com/yigithanbalci/golio/internal"
	"github.com/yigithanbalci/golio/internal/mock"
)

func TestChallengesClient_GetConfig(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*ChallengeConfigInfo
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*ChallengeConfigInfo{{}},
			doer: mock.NewJSONMockDoer([]ChallengeConfigInfo{{}}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&ChallengesClient{c: client}).GetConfig()
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestChallengesClient_GetPercentiles(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    PercentilesByChallenges
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: PercentilesByChallenges{},
			doer: mock.NewJSONMockDoer(PercentilesByChallenges{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&ChallengesClient{c: client}).GetPercentiles()
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestChallengesClient_GetConfigWithChallengeId(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *ChallengeConfigInfo
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &ChallengeConfigInfo{},
			doer: mock.NewJSONMockDoer(ChallengeConfigInfo{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&ChallengesClient{c: client}).GetConfigByChallengeID(1)
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestChallengesClient_GetLeaderBoardByChallengeIdAndLevel(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*ApexPlayerInfo
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*ApexPlayerInfo{{}},
			doer: mock.NewJSONMockDoer([]ApexPlayerInfo{{}}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&ChallengesClient{c: client}).GetLeaderBoardByChallengeIDAndLevel(203102, TierMaster, 15)
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestChallengesClient_GetPercentilesWithChallengeId(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    Percentiles
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: Percentiles{},
			doer: mock.NewJSONMockDoer(Percentiles{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&ChallengesClient{c: client}).GetPercentilesByChallengeID(1)
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}

func TestChallengesClient_GetPlayerDataWithPUUID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *PlayerInfo
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &PlayerInfo{},
			doer: mock.NewJSONMockDoer(PlayerInfo{}, 200),
		},
		{
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&ChallengesClient{c: client}).GetPlayerDataByPUUID("1")
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}
