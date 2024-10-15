package main

import (
	"testing"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is a test",
	}
	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		if err != nil {
			return
		}

		want := "this is a test"
		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		want := ErrNotFound

		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		d := Dictionary{}
		d.Add("test", "this is a test definition")

		want := "this is a test definition"
		got, err := d.Search("test")
		if err != nil {
			t.Fatal("should find added word: ", err)
		}
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, definition)
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	// expecting word not found
	assertError(t, err, ErrNotFound)
}
