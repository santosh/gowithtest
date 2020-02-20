package racer

import (
	"net/http"
	"time"
)

// Racer runs http.Get() on both passed URLs and returns the one responded first.
func Racer(url1, url2 string) (winner string) {
	startA := time.Now()
	http.Get(url1)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(url2)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return url1
	}
	return url2
}
