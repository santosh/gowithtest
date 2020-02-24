// Package sync demonstrates supplement for concurrency. WaitGroups, Mutexes.
package sync

// Counter data structure has a value field.
// value incerments with Inc()
type Counter struct {
	value int
}

// Inc increments the Counter.value by 1
func (c *Counter) Inc() {
	c.value++
}

// Value returns Counter.value
func (c *Counter) Value() int {
	return c.value
}
