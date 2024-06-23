package internal

import (
	"fmt"
	"net/http"
)

type RequestOption func(r *http.Request)

func WithHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		fmt.Println(value)
		r.Header.Add(key, value)
	}
}
