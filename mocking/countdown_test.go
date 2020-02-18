package mocking

import (
	"bytes"
	"io"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(io.Writer(buffer), 3)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
