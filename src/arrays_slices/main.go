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
	// declaring and setting array in one line
	var numbers = [5]int{5, 4, 3, 2, 1}
	// printing it
	fmt.Println(numbers)
}
