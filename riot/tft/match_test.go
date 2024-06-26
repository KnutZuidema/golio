package tft

import (
    "fmt"
    "github.com/KnutZuidema/golio/api"
    "github.com/KnutZuidema/golio/internal"
    "github.com/KnutZuidema/golio/internal/mock"
    "github.com/sirupsen/logrus"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "net/http"
    "testing"
)

func TestTFTMatch_GetMatchesByPUUID(t *testing.T) {
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
        t.Run(
            tt.name, func(t *testing.T) {
                client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
                got, err := (&MatchClient{c: client}).GetMatchesByPUUID("puuid")
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}

func TestTFTMatch_GetMatchByMatchID(t *testing.T) {
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
            doer: mock.NewJSONMockDoer(&Match{}, 200),
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
                got, err := (&MatchClient{c: client}).GetMatchByMatchID("1234")
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}
