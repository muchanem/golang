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
	// 1 is the second value, and we ar cutting everything below that value off so it is removing red
	colours = append(colours[1:len(colours)])
	fmt.Println(colours)
	// defualt is len of colours, just proving
	colours = append(colours[1:])
	fmt.Println(colours)
}
