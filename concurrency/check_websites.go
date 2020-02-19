// Package concurrency demonstrate optimization with concurrency using go routines.
//
// go routines are thread managed by go virtual machine (and are not system threads).
package concurrency

// WebsiteChecker is used in CheckWebsites.
type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

// CheckWebsites returns a map of each URL checked to a boolean value
// true for a good response, false for a bad response.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
