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
	// - to sepeate will be out putted
	// %b for binary gets applied to second number
	// \n an escape character for newline (println already does this hence ln)
	fmt.Printf("%d %b \n", 13, 13)
}
