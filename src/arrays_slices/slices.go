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
	// implicitly declaring slice and using make function to set type to int, intital lenth to 5 and the most values to 5
	numbers := make([]int, 5, 5)
	// seting the values
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 5
	numbers[3] = 7
	numbers[4] = 11
	fmt.Println(numbers)
	// adding number to slice
	numbers = append(numbers, 13)
	// print
	fmt.Println(numbers)
	// since 13 exaeded the cap it doubled
	fmt.Println(cap(numbers))

}
