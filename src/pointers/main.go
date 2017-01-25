package main

import "fmt"

func main() {
	// delcaring a pointer that points at a int value
	var pointer1 *int

	// checking is p  = nil (it does)
	if pointer1 != nil {
		// use asterisek to show it is pointer
		fmt.Println("The value of 'pointer' is:", *pointer1)
		// no asterisek for adress
		fmt.Println("The adress is:", pointer1)
	} else {
		fmt.Println("Pointer is nil")
	}
	value1 := 13
	pointer1 = &value1
	if pointer1 != nil {
		// use asterisek to show it is pointer
		fmt.Println("The value of pointer is:", *pointer1)
		// no asterisek for adress
		fmt.Println("The adress is:", pointer1)
	} else {
		fmt.Println("Pointer is nil")
	}
	value2 := 13.2
	pointer2 := &value2
	fmt.Println("Value of pointer2", *pointer2)
	*pointer2 = *pointer2 / 2
	fmt.Println("Pointer2:", *pointer2)
	fmt.Println("Value:", value2)

}
