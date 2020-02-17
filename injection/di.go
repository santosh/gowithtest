// Package injection demonstrate dependency injection.
package injection

import (
	"fmt"
	"io"
	"net/http"
)

// Greet ("santosh") -> "Hello, santosh"
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func myGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":8500", http.HandlerFunc(myGreeterHandler))
}
