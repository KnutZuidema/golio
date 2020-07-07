package internal

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal/mock"
)

func TestClient_DoRequest(t *testing.T) {
	t.Parallel()
	type args struct {
		method   string
		endpoint string
		body     io.Reader
	}
	tests := []struct {
		name    string
		args    args
		doer    Doer
		wantErr bool
	}{
		{
			name: "invalid method",
			args: args{
				method: "a\nb c",
			},
			doer:    mock.NewStatusMockDoer(200),
			wantErr: true,
		},
		{
			name: "error doer",
			args: args{
				method: "GET",
			},
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name: "fail second try",
			args: args{
				method: "GET",
			},
			doer:    failOnSecondDoer(),
			wantErr: true,
		},
		{
			name: "invalid retry header",
			args: args{
				method: "GET",
			},
			doer: mock.NewHeaderMockDoer(http.StatusTooManyRequests, http.Header{
				"Retry-After": []string{"abc"},
			}),
			wantErr: true,
		},
		{
			name: "invalid status",
			args: args{
				method: "GET",
			},
			doer:    mock.NewStatusMockDoer(199),
			wantErr: true,
		},
		{
			name: "rate limited",
			args: args{
				method: "GET",
			},
			doer: mock.NewRateLimitDoer(1),
		},
		{
			name: "service unavailable",
			args: args{
				method: "GET",
			},
			doer: mock.NewUnavailableOnceDoer(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionEuropeNorthEast, "", tt.doer, logrus.StandardLogger())
			_, err := c.DoRequest(tt.args.method, tt.args.endpoint, tt.args.body)
			assert.Equal(t, err != nil, tt.wantErr)
		})
	}
}

func TestClient_GetInto(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		doer    Doer
		wantErr bool
	}{
		{
			name:    "fail decode",
			target:  struct{}{},
			doer:    mock.NewJSONMockDoer(0, 200),
			wantErr: true,
		},
		{
			name: "fail request",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name:   "succeed",
			target: new(int),
			doer:   mock.NewJSONMockDoer(1, 200),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", tt.doer, logrus.StandardLogger())
			err := c.GetInto("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestClient_PostInto(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		doer    Doer
		wantErr bool
	}{
		{
			name:    "fail decode",
			target:  struct{}{},
			doer:    mock.NewJSONMockDoer(0, 200),
			wantErr: true,
		},
		{
			name: "fail request",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name:   "succeed",
			target: new(int),
			doer:   mock.NewJSONMockDoer(1, 200),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", tt.doer, logrus.StandardLogger())
			err := c.PostInto("endpoint", struct{}{}, tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestClient_Post(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		doer    Doer
		wantErr bool
	}{
		{
			name:    "fail encode",
			target:  mock.FailJSONEncoding{},
			doer:    mock.NewJSONMockDoer(0, 200),
			wantErr: true,
		},
		{
			name: "fail request",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", tt.doer, logrus.StandardLogger())
			_, err := c.Post("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestClient_Put(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		doer    Doer
		wantErr bool
	}{
		{
			name:    "fail encode",
			target:  mock.FailJSONEncoding{},
			doer:    mock.NewJSONMockDoer(0, 200),
			wantErr: true,
		},
		{
			name: "fail request",
			doer: &mock.Doer{
				Custom: func(r *http.Request) (*http.Response, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", tt.doer, logrus.StandardLogger())
			err := c.Put("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func failOnSecondDoer() Doer {
	count := 0
	return &mock.Doer{
		Custom: func(r *http.Request) (*http.Response, error) {
			if count == 0 {
				count++
				return mock.NewStatusMockDoer(http.StatusServiceUnavailable).Do(r)
			}
			return nil, fmt.Errorf("error")
		},
	}
}
