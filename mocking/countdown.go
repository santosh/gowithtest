// Package mocking makes use of dependency injection to fake
// time.Sleep.
package mocking

import (
	"fmt"
	"io"
	"time"
)

// Countdown counts down from 'top' to 1, and prints 'Go!'
// in place of 0. Writing is done to the io.Writer
func Countdown(w io.Writer, upper int) {
	for i := upper; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(w, "Go!")
}
