package val

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

func TestChallengesClient_GetPlatformData(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    *PlatformDataDto
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: &PlatformDataDto{},
			doer: mock.NewJSONMockDoer(PlatformDataDto{}, 200),
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
