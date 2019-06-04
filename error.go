package golio

import (
	"net/http"
)

type Error struct {
	Message    string
	StatusCode int
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrBadRequest = Error{
		Message:    "bad request",
		StatusCode: http.StatusBadRequest,
	}
	ErrUnauthorized = Error{
		Message:    "unauthorized",
		StatusCode: http.StatusUnauthorized,
	}
	ErrForbidden = Error{
		Message:    "forbidden",
		StatusCode: http.StatusForbidden,
	}
	ErrNotFound = Error{
		Message:    "not found",
		StatusCode: http.StatusNotFound,
	}
	ErrMethodNotAllowed = Error{
		Message:    "method not allowed",
		StatusCode: http.StatusMethodNotAllowed,
	}
	ErrUnsupportedMediaType = Error{
		Message:    "unsupported media type",
		StatusCode: http.StatusUnsupportedMediaType,
	}
	ErrRateLimitExceeded = Error{
		Message:    "rate limit exceeded",
		StatusCode: http.StatusTooManyRequests,
	}
	ErrInternalServerError = Error{
		Message:    "internal server error",
		StatusCode: http.StatusInternalServerError,
	}
	ErrBadGateway = Error{
		Message:    "bad gateway",
		StatusCode: http.StatusBadGateway,
	}
	ErrServiceUnavailable = Error{
		Message:    "service unavailable",
		StatusCode: http.StatusServiceUnavailable,
	}
	ErrGatewayTimeout = Error{
		Message:    "gateway timeout",
		StatusCode: http.StatusGatewayTimeout,
	}
	StatusToError = map[int]Error{
		http.StatusBadRequest:           ErrBadRequest,
		http.StatusUnauthorized:         ErrUnauthorized,
		http.StatusForbidden:            ErrForbidden,
		http.StatusNotFound:             ErrNotFound,
		http.StatusMethodNotAllowed:     ErrMethodNotAllowed,
		http.StatusUnsupportedMediaType: ErrUnsupportedMediaType,
		http.StatusTooManyRequests:      ErrRateLimitExceeded,
		http.StatusInternalServerError:  ErrInternalServerError,
		http.StatusBadGateway:           ErrBadGateway,
		http.StatusServiceUnavailable:   ErrServiceUnavailable,
		http.StatusGatewayTimeout:       ErrGatewayTimeout,
	}
)
