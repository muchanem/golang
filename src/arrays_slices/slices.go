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
	// saying give me back all items ecxept the first one (you dont need length of colurs, that is already the default)
	colours = append(colours[1:len(colours)])
	fmt.Println(colours)
	// removing last item
	colours = append(colours[:len(colours)-1])
	fmt.Println(colours)
}
