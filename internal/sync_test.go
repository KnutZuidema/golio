package internal

import (
	"sync"
	"testing"
)

func TestRWLockToggle(t *testing.T) {
	var mu sync.RWMutex
	unlock, _ := RWLockToggle(&mu)
	unlock()
	unlock, toggle := RWLockToggle(&mu)
	toggle()
	unlock()
}
