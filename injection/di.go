// Package injection demonstrate dependency injection.
package injection

import (
	"bytes"
	"fmt"
)

// Greet ("santosh") -> "Hello, santosh"
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
