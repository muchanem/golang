package main

/*
  Printing numbers in hex
*/
import (
	"fmt"
)

// using prinf to print with formating
// %d for decimal gets applied to first number
// %b for binary gets applied to second number
// %x prints in hex with lowercase letters
// \n an escape character for newline (println already does this hence ln)
func main() {
	fmt.Printf("%d %b %x \n", 13, 13, 13)
}
