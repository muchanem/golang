package main

import "fmt"

func main() {
// same as number loop just using %q which prints the value in utf8
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n ", i, i, i, i)

	}

}
