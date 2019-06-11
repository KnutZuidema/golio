package mock

import (
	"net/http"
	"testing"
	"time"
)

func TestNewJSONMockDoer(t *testing.T) {
	type args struct {
		object interface{}
		code   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "contruct",
			args: args{
				object: struct {
					Attr int
				}{
					Attr: 1,
				},
				code: 200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewJSONMockDoer(tt.args.object, tt.args.code)
			if got == nil {
				t.Error("returned nil")
			}
		})
	}
}

func TestNewStatusMockDoer(t *testing.T) {
	type args struct {
		code int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "contruct",
			args: args{
				code: 200,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStatusMockDoer(tt.args.code); got == nil {
				t.Error("returned nil")
			}
		})
	}
}

func TestNewHeaderMockDoer(t *testing.T) {
	type args struct {
		code    int
		headers http.Header
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "construct",
			args: args{
				code: 200,
				headers: http.Header{
					"key": []string{"value"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeaderMockDoer(tt.args.code, tt.args.headers); got == nil {
				t.Error("returned nil")
			}
		})
	}
}

func TestDoer_Do(t *testing.T) {
	tests := []struct {
		name    string
		wait    time.Duration
		custom  func(r *http.Request) (*http.Response, error)
		wantErr bool
	}{
		{
			name: "do request",
			wait: time.Second,
		},
		{
			name: "custom",
			custom: func(r *http.Request) (*http.Response, error) {
				return &http.Response{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NewStatusMockDoer(200)
			d.Custom = tt.custom
			d.ResponseTime = tt.wait
			before := time.Now()
			got, err := d.Do(&http.Request{})
			after := time.Now()
			if after.Sub(before) < tt.wait {
				t.Error("did not wait specified duration")
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Doer.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.custom != nil && got == nil {
				t.Error("custom function returned nil")
				return
			} else if tt.custom == nil && got != &d.Response {
				t.Errorf("Doer.Do() = %v, want %v", got, d.Response)
			}
		})
	}
}
