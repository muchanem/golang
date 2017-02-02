package main

import (
	"fmt"
)

func main() {
	// declaring and setting a sclice

	var colours = []string{"Red", "Orange", "Green"}
	fmt.Println(colours)
	// adding another colour using append
	colours = append(colours, "Blue")
	fmt.Println(colours)
}
