package internal

import (
	"net/http"
)

// Doer is an interface for any client that can process an HTTP request and return a response.
// This will most commonly be a simple HTTP client.
type Doer interface {
	// Do processes an HTTP request and returns the response
	Do(r *http.Request) (*http.Response, error)
}
