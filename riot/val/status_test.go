package val

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

func TestChallengesClient_GetPlatformData(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *PlatformData
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &PlatformData{},
			doer: mock.NewJSONMockDoer(PlatformData{}, 200),
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
			got, err := (&StatusClient{c: client}).GetPlatformData()
			require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
			if tt.wantErr == nil {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
