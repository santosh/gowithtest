// Package mocking makes use of dependency injection to fake
// time.Sleep.
package mocking

import (
	"fmt"
	"io"
	"time"
)

// Sleeper interface gives independence to use different Sleeper
// in different situations. For example, we don't want to time.Sleep
// in tests.
type Sleeper interface {
	Sleep(n int)
}

// SpySleeper has a Calls variable which increments everytime
// SpySleeper.Sleep() is called. Better for use in tests.
type SpySleeper struct {
	Calls int
}

// Sleep spies on number of Sleep calls. It maintains
// a variable which is incremented each time it is called.
func (s *SpySleeper) Sleep(n int) {
	s.Calls++
}

// DefaultSleeper implements replacement for time.Sleep.
// Because we can't withhold giving up execution time for
// Sleep on tests.
type DefaultSleeper struct{}

// Sleep will pause execution for n seconds.
func (d *DefaultSleeper) Sleep(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}

// Countdown counts down from 'top' to 1, and prints 'Go!'
// in place of 0. Writing is done to the io.Writer
func Countdown(w io.Writer, upper int, sleeper Sleeper) {
	for i := upper; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep(1)
	}
	fmt.Fprint(w, "Go!")
}
