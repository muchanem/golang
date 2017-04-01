package main

/*
  Printing numbers in binary
  See godoc.org/fmt for more verbs
*/
import (
	"fmt"
)

func main() {
	// using prinf to print with formating
	// %d for decimal gets applied to first number
	// - to sepeate
	// %b for binary gets applied to second number
	fmt.Printf("%d - %b", 13, 13)
}
