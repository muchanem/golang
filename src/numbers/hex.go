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
// %#x prints in hex with lowercase letters with a zero x in front
// %x prints in hex lowercase letters
// you can change to capital letters by doing %X
// \n an escape character for newline (println already does this hence ln)
func main() {
	fmt.Printf("%d \t %b \t %#x - %x \n", 13, 13, 13, 13)
}
