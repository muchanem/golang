package main

import (
	"fmt"
)

func main() {
	// setting array of strings called colours
	var colours [3]string
	colours[0] = "Red"
	colours[1] = "Orange"
	colours[2] = "Yellow"
	// printing array
	fmt.Println(colours)
	// use this method to get to single item in array
	fmt.Println(colours[2])
}
