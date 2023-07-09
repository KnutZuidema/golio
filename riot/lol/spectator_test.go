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

func TestSpectatorClient_ListFeatured(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *FeaturedGames
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &FeaturedGames{},
			doer: mock.NewJSONMockDoer(FeaturedGames{}, 200),
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
			got, err := (&SpectatorClient{c: client}).ListFeatured()
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestSpectatorClient_GetCurrent(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *GameInfo
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &GameInfo{},
			doer: mock.NewJSONMockDoer(GameInfo{}, 200),
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
			got, err := (&SpectatorClient{c: client}).GetCurrent("id")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
