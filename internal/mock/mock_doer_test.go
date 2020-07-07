package mock

import (
	"encoding/json"
	"fmt"
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
			name: "construct",
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

func ExampleNewJSONMockDoer() {
	type object struct {
		Attr int
	}
	input := object{
		Attr: 2,
	}
	doer := NewJSONMockDoer(input, 200)
	request, _ := http.NewRequest("GET", "https://example.com", nil)
	response, _ := doer.Do(request)
	var output object
	_ = json.NewDecoder(response.Body).Decode(&output)
	fmt.Printf("status code: %d, body: %+v\n", response.StatusCode, output)
	// Output: status code: 200, body: {Attr:2}
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
			name: "construct",
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

func ExampleNewStatusMockDoer() {
	doer := NewStatusMockDoer(200)
	request, _ := http.NewRequest("GET", "http://example.com", nil)
	response, _ := doer.Do(request)
	fmt.Printf("status code: %d", response.StatusCode)
	// Output: status code: 200
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

func ExampleNewHeaderMockDoer() {
	doer := NewHeaderMockDoer(200, http.Header{"Content-Type": []string{"application/json"}})
	request, _ := http.NewRequest("GET", "http://example.com", nil)
	response, _ := doer.Do(request)
	fmt.Printf("status code: %d, content-type: %v", response.StatusCode, response.Header.Get("Content-Type"))
	// Output: status code: 200, content-type: application/json
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

func TestFailJSONEncoding_MarshalJSON(t *testing.T) {
	if _, err := json.Marshal(FailJSONEncoding{}); err == nil {
		t.Errorf("should return error")
	}
}

func TestNewRateLimitDoer(t *testing.T) {
	type args struct {
		object interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "construct",
			args: args{
				object: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRateLimitDoer(tt.args.object)
			if got == nil {
				t.Fatal("returned nil")
			}
			if r, err := got.Do(nil); err != nil {
				t.Error(err)
			} else if r.StatusCode != http.StatusTooManyRequests {
				t.Errorf("expected status %d, got %d", http.StatusTooManyRequests, r.StatusCode)
			}
			if r, err := got.Do(nil); err != nil {
				t.Error(err)
			} else if r.StatusCode != http.StatusOK {
				t.Errorf("expected status %d, got %d", http.StatusOK, r.StatusCode)
			}
		})
	}
}

func TestNewUnavailableOnceDoer(t *testing.T) {
	type args struct {
		object interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "construct",
			args: args{
				object: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUnavailableOnceDoer(tt.args.object)
			if got == nil {
				t.Fatal("returned nil")
			}
			if r, err := got.Do(nil); err != nil {
				t.Error(err)
			} else if r.StatusCode != http.StatusServiceUnavailable {
				t.Errorf("expected status %d, got %d", http.StatusServiceUnavailable, r.StatusCode)
			}
			if r, err := got.Do(nil); err != nil {
				t.Error(err)
			} else if r.StatusCode != http.StatusOK {
				t.Errorf("expected status %d, got %d", http.StatusOK, r.StatusCode)
			}
		})
	}
}
