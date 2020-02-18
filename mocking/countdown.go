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

// DefaultSleeper implements replacement for time.Sleep.
// Because we can't withhold giving up execution time for
// Sleep on tests.
type DefaultSleeper struct{}

// Sleep will pause execution for n seconds.
func (d *DefaultSleeper) Sleep(n int) {
	time.Sleep(time.Duration(n) * time.Second)
}

// CountdownOperationSpy keep records of calls to its methods
type CountdownOperationSpy struct {
	Calls []string
}

// Sleep records sleep call on CountdownOperationSpy.Calls
func (s *CountdownOperationSpy) Sleep(n int) {
	s.Calls = append(s.Calls, sleep)
}

// Write records write call on CountdownOperationSpy.Calls
func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// Countdown counts down from 'top' to 1, and prints 'Go!'
// in place of 0. Writing is done to the io.Writer
func Countdown(w io.Writer, upper int, sleeper Sleeper) {
	for i := upper; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep(1)
	}
	fmt.Fprint(w, "Go!")
}
