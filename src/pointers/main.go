package main

import "fmt"

func main() {
	// delcaring a pointer that points at a int value
	var p *int

	// checking is p  = nil (it does)
	if p != nil {
		// use asterisek to show it is pointer
		fmt.Println("The value of p is:", *p)
		// no asterisek for adress
		fmt.Println("The adress is:", p)
	} else {
		fmt.Println("P is nil")
	}
	v := 13
	p = &v
	if p != nil {
		// use asterisek to show it is pointer
		fmt.Println("The value of p is:", *p)
		// no asterisek for adress
		fmt.Println("The adress is:", p)
	} else {
		fmt.Println("P is nil")
	}
}
