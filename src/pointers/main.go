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
	// defining another value
	value2 := 13.2
	// setting pointer2 to that value
	pointer2 := &value2
	// printing the value of what the pointer is pointing at
	fmt.Println("Value of pointer2", *pointer2)
	// dividing the value that the pointer is pointing at by 2
	*pointer2 = *pointer2 / 2
	// printing the value of what the pointer is pointing at
	fmt.Println("Pointer2:", *pointer2)
	// showing that my divide by two operation actually effected the pointer
	fmt.Println("Value:", value2)

}
