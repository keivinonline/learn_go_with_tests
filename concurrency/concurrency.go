package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// // Blocking operation
		// results[url] = wc(url)

		// // Goroutine via anonymous functions
		// // each iteration of the loop starts a new goroutine
		go func(u string) {
			// // Send statement to the channel
			// // channel <- value
			resultChannel <- result{u, wc(u)}
		}(url)
	}
	// Once all goroutines are done executing
	// // fetch the results in the channel and fill up the results map
	for i := 0; i < len(urls); i++ {
		// // Receive statement
		// // value := <- channel
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
