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

func TestTFTLeague_GetChallenger(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name    string
        want    *LeagueList
        doer    internal.Doer
        wantErr error
    }{
        {
            name: "get response",
            want: &LeagueList{},
            doer: mock.NewJSONMockDoer(LeagueList{}, 200),
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
                got, err := (&LeagueClient{c: client}).GetChallenger(QueueRankedTFT)
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}

func TestTFTLeague_GetEntriesBySummoner(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name    string
        want    []*LeagueEntry
        doer    internal.Doer
        wantErr error
    }{
        {
            name: "get response",
            want: []*LeagueEntry{},
            doer: mock.NewJSONMockDoer([]*LeagueEntry{}, 200),
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
                got, err := (&LeagueClient{c: client}).GetEntriesBySummoner("summonerId")
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}

func TestTFTLeague_GetEntries(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name    string
        want    []*LeagueEntry
        doer    internal.Doer
        wantErr error
    }{
        {
            name: "get response",
            want: []*LeagueEntry{},
            doer: mock.NewJSONMockDoer([]*LeagueEntry{}, 200),
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
                got, err := (&LeagueClient{c: client}).GetEntries("DIAMOND", "I")
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}

func TestTFTLeague_GetGrandMaster(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name    string
        want    *LeagueList
        doer    internal.Doer
        wantErr error
    }{
        {
            name: "get response",
            want: &LeagueList{},
            doer: mock.NewJSONMockDoer(&LeagueList{}, 200),
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
                got, err := (&LeagueClient{c: client}).GetGrandMaster(QueueRankedTFT)
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}

func TestTFTLeague_GetLeagues(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name    string
        want    *LeagueList
        doer    internal.Doer
        wantErr error
    }{
        {
            name: "get response",
            want: &LeagueList{},
            doer: mock.NewJSONMockDoer(&LeagueList{}, 200),
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
                got, err := (&LeagueClient{c: client}).GetLeagues("1234")
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}

func TestTFTLeague_GetMaster(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name    string
        want    *LeagueList
        doer    internal.Doer
        wantErr error
    }{
        {
            name: "get response",
            want: &LeagueList{},
            doer: mock.NewJSONMockDoer(&LeagueList{}, 200),
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
                got, err := (&LeagueClient{c: client}).GetMaster(QueueRankedTFT)
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}

func TestTFTLeague_GetRatedLaddersByQueue(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name    string
        want    []*TopRatedLadderEntry
        doer    internal.Doer
        wantErr error
    }{
        {
            name: "get response",
            want: []*TopRatedLadderEntry{},
            doer: mock.NewJSONMockDoer([]*TopRatedLadderEntry{}, 200),
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
                got, err := (&LeagueClient{c: client}).GetRatedLaddersByQueue(QueueRankedTFT)
                require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
                if tt.wantErr == nil {
                    assert.Equal(t, got, tt.want)
                }
            },
        )
    }
}
