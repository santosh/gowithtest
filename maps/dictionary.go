// Package maps module demonstrate usage of map data structure.
//
// In this module we build a dictionay.
package maps

const (
	// ErrNotFound should be raised in case there is no word in dictionary
	ErrNotFound = DictionaryErr("cannot find the word you were looking for")

	// ErrWordExists should be raised when adding a new entry to dictionary
	// and entry already exists
	ErrWordExists = DictionaryErr("cannot add the word because it already exists")

	// ErrWordNotExists is typically raised when trying to Update existing
	// definitions.
	ErrWordNotExists = DictionaryErr("cannot update the word because it doesn't exists")
)

// Dictionary represent a map of word and meaning
type Dictionary map[string]string

// DictionaryErr and subsequent func definition fulfills error interface
type DictionaryErr string

// Error implements error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

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

// Update updates the value of given word (key).
// If there is no word already, it doesn't adds.
func (d Dictionary) Update(word, meaning string) error {
	_, ok := d[word]
	if ok {
		d[word] = meaning
		return nil
	}

	return ErrWordNotExists
}
