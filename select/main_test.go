package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				time.Sleep(delay)
				w.WriteHeader(http.StatusOK)
			}))
}

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers, retruning rul of fastest one",
		func(t *testing.T) {

			slowServer := makeDelayedServer(20 * time.Millisecond)
			fastServer := makeDelayedServer(0 * time.Millisecond)
			// calls the 2 following `at end of containing function`
			defer slowServer.Close()
			defer fastServer.Close()

			slowUrl := slowServer.URL
			fastUrl := fastServer.URL

			want := fastUrl
			got, _ := Racer(slowUrl, fastUrl)

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	// t.Run("returns an error if a server does not respond within 10 seconds",
	// 	func(t *testing.T) {
	// 		serverA := makeDelayedServer(11 * time.Second)
	// 		serverB := makeDelayedServer(11 * time.Second)

	// 		defer serverA.Close()
	// 		defer serverB.Close()

	// 		_, err := Racer(serverA.URL, serverB.URL)
	// 		// we are expecting timeout error and not nil
	// 		if err == nil {
	// 			t.Error("expected an error but didn't get any")
	// 		}
	// 	})
	t.Run("returns an error if server does not respond within timeout", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error byt didn't get one")
		}
	})
}
