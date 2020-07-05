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
			doer:    invalidHeaderDoer(),
			wantErr: true,
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
		wantErr bool
	}{
		{
			name:    "fail decode",
			target:  struct{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", mock.NewJSONMockDoer(0, 200), logrus.StandardLogger())
			err := c.GetInto("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestClient_PostInto(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "fail decode",
			target:  struct{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", mock.NewJSONMockDoer(0, 200), logrus.StandardLogger())
			err := c.PostInto("endpoint", struct{}{}, tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestClient_post(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "fail encode",
			target:  mock.FailJSONEncoding{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", mock.NewStatusMockDoer(200), logrus.StandardLogger())
			_, err := c.Post("endpoint", tt.target)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestClient_put(t *testing.T) {
	tests := []struct {
		name    string
		target  interface{}
		wantErr bool
	}{
		{
			name:    "fail encode",
			target:  mock.FailJSONEncoding{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(api.RegionOceania, "API_KEY", mock.NewStatusMockDoer(200), logrus.StandardLogger())
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

func invalidHeaderDoer() Doer {
	return mock.NewHeaderMockDoer(http.StatusTooManyRequests, http.Header{
		"Retry-After": []string{"abc"},
	})
}
