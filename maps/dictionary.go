// Package maps module demonstrate usage of map data structure.
//
// In this module we build a dictionay.
package maps

import "errors"

var (
	// ErrNotFound should be raised in case there is no word in dictionary
	ErrNotFound = errors.New("could not find the word you were looking for")
	// ErrWordExists should be raised when adding a new entry to dictionary
	// and entry already exists
	ErrWordExists = errors.New("cannot add the word because it already exists")
)

// Dictionary represent a map of word and meaning
type Dictionary map[string]string

// Search iterates over a map to find word, then returns it's value,
// which is represent the meaning
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add inserts a word and a meaning to the caller.
func (d Dictionary) Add(word, meaning string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}
	d[word] = meaning
	return nil
}
