// Package sync demonstrates supplement for concurrency. WaitGroups, Mutexes.
package sync

import "sync"

// Counter data structure has a value field.
// value incerments with Inc()
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments the Counter.value by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns Counter.value
func (c *Counter) Value() int {
	return c.value
}
