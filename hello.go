package main

import "fmt"

// Hello returns "Hello, world"
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
