// Package maps module demonstrate usage of map data structure.
//
// In this module we build a dictionay.
package maps

import "errors"

// Search iterates over a map to find word, then returns it's value,
// which is represent the meaning
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}

type Dictionary map[string]string
