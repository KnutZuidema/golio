package tft

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

func TestTournamentClient_CreateCodes(t *testing.T) {
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
			got, err := (&TournamentClient{c: client}).CreateCodes(0, 0, &TournamentCodeParameters{}, true)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestTournamentClient_ListLobbyEvents(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *LobbyEventList
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &LobbyEventList{},
			doer: mock.NewJSONMockDoer(LobbyEventList{}, 200),
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
			got, err := (&TournamentClient{c: client}).ListLobbyEvents("code", true)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestTournamentClient_CreateProvider(t *testing.T) {
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
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&TournamentClient{c: client}).CreateProvider(&ProviderRegistrationParameters{}, true)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestTournamentClient_Create(t *testing.T) {
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
			name:    "not found",
			wantErr: api.ErrNotFound,
			doer:    mock.NewStatusMockDoer(http.StatusNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
			got, err := (&TournamentClient{c: client}).Create(&TournamentRegistrationParameters{}, true)
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestTournamentClient_Get(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *Tournament
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &Tournament{},
			doer: mock.NewJSONMockDoer(Tournament{}, 200),
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
			got, err := (&TournamentClient{c: client}).Get("code")
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}

func TestTournamentClient_Update(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			doer: mock.NewStatusMockDoer(200),
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
			err := (&TournamentClient{c: client}).Update("code", TournamentUpdateParameters{})
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
		})
	}
}
