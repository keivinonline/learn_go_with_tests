package main

import "errors"

type Dictionary map[string]string

var (
	ErrUnknownWord = errors.New("unknown word")
	ErrNotFound    = errors.New("word not found")
	ErrWordExists  = errors.New("word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	if d[word] != "" {
		return d[word], nil
	}
	return "", ErrNotFound
}
func (d Dictionary) Add(word, definition string) error {
	// checks for existing word
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil: // no err means word is found
		return ErrWordExists
	default:
		return err // catch all
	}

	if d[word] != "" {
		return ErrWordExists
	} else {
		d[word] = definition
		return nil
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
