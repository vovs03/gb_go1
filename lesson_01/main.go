package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println
	// g = greeting
	var g = "Hello, Gophers!"
	p(g, s.Count(g, "o"))
}
