package mock

import (
	"bytes"
	"io"
	"testing"
)

func TestResponseBody_Read(t *testing.T) {
	tests := []struct {
		name      string
		content   []byte
		read      []byte
		want      []byte
		wantN     int
		wantErr   error
		closeBody bool
	}{
		{
			name:    "read partial",
			content: []byte("content"),
			read:    make([]byte, 4),
			want:    []byte("cont"),
			wantN:   4,
		},
		{
			name:    "read full",
			content: []byte("content"),
			read:    make([]byte, 7),
			want:    []byte("content"),
			wantN:   7,
			wantErr: io.EOF,
		},
		{
			name:    "read big target",
			content: []byte("content"),
			read:    make([]byte, 10),
			want:    []byte("content"),
			wantN:   7,
			wantErr: io.EOF,
		},
		{
			name:    "read empty",
			read:    make([]byte, 10),
			wantN:   0,
			wantErr: io.EOF,
		},
		{
			name:      "read closed",
			wantErr:   ErrBodyClosed,
			closeBody: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ResponseBody{
				Content: tt.content,
				closed:  tt.closeBody,
			}
			gotN, err := b.Read(tt.read)
			if err != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("got %v bytes read, want %v", gotN, tt.wantN)
			}
			if !bytes.Equal(tt.want, tt.read[:tt.wantN]) {
				t.Errorf("got %s, want %s", string(tt.read), string(tt.want))
			}
		})
	}
}

func TestResponseBody_Close(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "close",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ResponseBody{}
			if err := b.Close(); err != nil {
				t.Error("err is not nil")
			}
			if !b.closed {
				t.Error("flag not set to true")
			}
		})
	}
}
