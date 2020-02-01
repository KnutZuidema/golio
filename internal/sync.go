package internal

import "sync"

// RWLockToggle locks the given mutex for reading and returns two functions
// the first function returned should be used to unlock the mutex
// the second function returned will unlock the read lock and instead lock the mutex for writing
// the unlock function will always call the correct unlock method
func RWLockToggle(mu *sync.RWMutex) (func(), func()) {
	sw := true
	mu.RLock()
	return func() {
			if sw {
				mu.RUnlock()
			} else {
				mu.Unlock()
			}
		},
		func() {
			sw = false
			mu.RUnlock()
			mu.Lock()
		}
}
