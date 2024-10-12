package main

import "testing"

// testing.TB is an interface for T (testing) and B (Benchmark)
func assertCorrectMessage(t testing.TB, got, want string) {
	// Tells test suite this is only a helper
	// Report the function call rather than this helper
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

type Test struct {
	language string
	name     string
	want     string
}

var Tests = []Test{
	{language: "English", name: "John", want: englishHelloPrefix + "John"},
	{language: "Spanish", name: "Senor", want: spanishHelloPrefix + "Senor"},
	{language: "French", name: "Clement", want: frenchHelloPrefix + "Clement"},
}

var Tests2 = []Test{
	{language: "English", name: "", want: englishHelloPrefix + defaultName},
	{language: "Spanish", name: "", want: spanishHelloPrefix + defaultName},
	{language: "French", name: "", want: frenchHelloPrefix + defaultName},
}

func TestHello(t *testing.T) {
	for _, testCase := range Tests {
		t.Run(testCase.language, func(t *testing.T) {
			got := Hello(testCase.name, testCase.language)
			assertCorrectMessage(t, got, testCase.want)
		})
	}

	for _, testCase := range Tests {
		t.Run(testCase.language+defaultName, func(t *testing.T) {
			got := Hello(testCase.name, testCase.language)
			assertCorrectMessage(t, got, testCase.want)
		})
	}
}
