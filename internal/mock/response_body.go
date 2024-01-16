package mock

import (
	"fmt"
	"io"
)

var (
	// ErrBodyClosed is returned if read is attempted on a closed body
	ErrBodyClosed = fmt.Errorf("body is closed")
)

// ResponseBody implements the ReadCloser interface to act as a response body for testing purposes
type ResponseBody struct {
	Content []byte
	lastPos int
	closed  bool
}

// Read reads from the content of the response body into p
func (b *ResponseBody) Read(p []byte) (n int, err error) {
	if b.closed {
		return 0, ErrBodyClosed
	}
	if b.lastPos == len(b.Content)-1 {
		return 0, io.EOF
	}
	nextPos := b.lastPos + len(p)
	var err error
	if nextPos >= len(b.Content) {
		nextPos = len(b.Content) - 1
		err = io.EOF
	}
	copy(p, b.Content[b.lastPos:nextPos])
	read := nextPos - b.lastPos
	b.lastPos = nextPos
	return read, err
}

// Close closes the response body. After closing the response body it cannot be read from
func (b *ResponseBody) Close() error {
	b.closed = true
	return nil
}
