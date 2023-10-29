package lor

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

func TestRankedClient_GetMasters(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		want    []*Player
		doer    internal.Doer
		wantErr error
	}{
		{
			name: "get response",
			want: []*Player{},
			doer: mock.NewJSONMockDoer([]*Player{}, 200),
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
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				client := internal.NewClient(api.RegionEuropeWest, "API_KEY", tt.doer, logrus.StandardLogger())
				got, err := (&RankedClient{c: client}).GetMasters()
				require.Equal(t, err, tt.wantErr, fmt.Sprintf("want err %v, got %v", tt.wantErr, err))
				if tt.wantErr == nil {
					assert.Equal(t, got, tt.want)
				}
			},
		)
	}
}
