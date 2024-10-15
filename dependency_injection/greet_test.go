package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Buffer type implements Writer interface
	// as it has the Write() method
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kei")

	got := buffer.String()
	want := "Hello, Kei"
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
