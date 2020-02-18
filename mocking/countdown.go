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
	Sleep()
}

// ConfigurableSleeper is a wrapper around time.Sleep.
// It combines duration of at same place.
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep sleeps for configurable amount of time. duration comes
// from the receiver
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// SpyTime voyeurs on sleeping time. To be used by tests.
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep is a fake time.Sleep
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// CountdownOperationSpy keep records of calls to its methods
type CountdownOperationSpy struct {
	Calls []string
}

// Sleep records sleep call on CountdownOperationSpy.Calls
func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write records write call on CountdownOperationSpy.Calls
func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// Countdown counts down from 'upper' to 1, and prints 'Go!'
// in place of 0. Writing is done to the io.Writer
func Countdown(w io.Writer, upper int, sleeper Sleeper) {
	for i := upper; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprint(w, "Go!")
}
