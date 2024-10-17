package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// test func without making actual HTTP calls
func mockWebsiteChecker(url string) bool {
	return url != "waat://fakesite.com"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://yahoo.com",
		"waat://fakesite.com",
	}
	want := map[string]bool{
		"https://google.com": true,
		"https://yahoo.com":  true,
		"waat://fakesite.com": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v got %v", want, got)
	}
}

// input string is not used
// simulate slow valid site
func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	// Create a slice of 100 urls
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	// Reset the time of test before it runs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
