package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "https://atlas.bydeluxe.com/confluence"
	fastURL := "https://rocket.methodstudios.com/"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
