package mock

import (
	"encoding/json"
	"net/http"
	"time"
)

// Doer is an implementation of the Doer interface for testing purposes which always returns the set response
// attribute
type Doer struct {
	Response     http.Response
	ResponseBody *ResponseBody
	ResponseTime time.Duration
	Custom       func(r *http.Request) (*http.Response, error)
}

// NewJSONMockDoer constructs a new MockDoer with the json representation of an object as body and the given status code
// CAUTION: silently returns an empty response if object is fails json marshaling
func NewJSONMockDoer(object interface{}, code int) *Doer {
	buffer, _ := json.Marshal(object)
	body := ResponseBody{
		Content: buffer,
	}
	return &Doer{
		Response: http.Response{
			Body:       &body,
			StatusCode: code,
		},
		ResponseBody: &body,
	}
}

// NewStatusMockDoer constructs a new MockDoer with the given status code
func NewStatusMockDoer(code int) *Doer {
	return &Doer{
		Response: http.Response{
			StatusCode: code,
		},
	}
}

// NewHeaderMockDoer constructs a new MockDoer with the given headers and status code
func NewHeaderMockDoer(code int, headers http.Header) *Doer {
	return &Doer{
		Response: http.Response{
			StatusCode: code,
			Header:     headers,
		},
	}
}

// Do returns the set HTTP response attribute after waiting for the specified amount of time
// if a custom function is specified it is used instead
func (d *Doer) Do(r *http.Request) (*http.Response, error) {
	if d.Custom != nil {
		return d.Custom(r)
	}
	timer := time.NewTimer(d.ResponseTime)
	<-timer.C
	// reset response body state
	if d.ResponseBody != nil {
		d.ResponseBody.lastPos = 0
		d.ResponseBody.closed = false
		d.Response.Body = d.ResponseBody
	}
	return &d.Response, nil
}
