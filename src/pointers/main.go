package main

import "fmt"

func main() {
	// delcaring a pointer that points at a int value
	var pointer *int

	// checking is p  = nil (it does)
	if pointer != nil {
		// use asterisek to show it is pointer
		fmt.Println("The value of 'pointer' is:", *pointer)
		// no asterisek for adress
		fmt.Println("The adress is:", pointer)
	} else {
		fmt.Println("Pointer is nil")
	}
	variable1 := 13
	pointer = &variable1
	if pointer != nil {
		// use asterisek to show it is pointer
		fmt.Println("The value of pointer is:", *pointer)
		// no asterisek for adress
		fmt.Println("The adress is:", pointer)
	} else {
		fmt.Println("Pointer is nil")
	}

}
