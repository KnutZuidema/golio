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
	if b.lastPos < len(b.Content) && len(b.Content) <= b.lastPos+len(p) {
		copy(p, b.Content[b.lastPos:])
		read := len(b.Content) - b.lastPos
		b.lastPos += read
		return read, io.EOF
	} else if b.lastPos >= len(b.Content) {
		return 0, io.EOF
	}
	copy(p, b.Content[b.lastPos:b.lastPos+len(p)])
	b.lastPos += len(p)
	return len(p), nil
}

// Close closes the response body. After closing the response body it cannot be read from
func (b *ResponseBody) Close() error {
	b.closed = true
	return nil
}
