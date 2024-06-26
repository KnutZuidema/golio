package tft

import (
    "fmt"
    "net/http"
    "testing"

    "github.com/KnutZuidema/golio/api"
    "github.com/KnutZuidema/golio/internal"
    "github.com/KnutZuidema/golio/internal/mock"
    "github.com/sirupsen/logrus"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestTFTSpectator_GetActiveGamesByPUUID(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name    string
        want    *CurrentGameInfo
        doer    internal.Doer
        wantErr error
    }{
        {
            name: "get response",
            want: &CurrentGameInfo{},
            doer: mock.NewJSONMockDoer(CurrentGameInfo{}, 200),
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
                got, err := (&SpectatorClient{c: client}).GetActiveGamesByPUUID("puuid")
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}

func TestTFTSpectator_GetFeaturedGames(t *testing.T) {
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
        t.Run(
            tt.name, func(t *testing.T) {
                client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
                got, err := (&SpectatorClient{c: client}).GetFeaturedGames()
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}
