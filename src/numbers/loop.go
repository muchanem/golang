package main

import (
	"fmt"
)

/*
  Printing numbers in using a loop in different formats
*/

func main() {
	/*
			   Loops:
			     * the first thing is what is the intalizer of the first value the loop holds
			     * the second thing is under what condition the loop will run
		       * the third thing is what will happen after the loop runs
	*/
	for i := 0; i < 200; i++ {
		/*
		   using prinf to print with formating
		   %d for decimal gets applied to first number
		   %b for binary gets applied to second number
		   %#x prints in hex with lowercase letters with a zero x in front
		   %x prints in hex lowercase letters
		   you can change to capital letters by doing %X
		   \n an escape character for newline (println already does this hence ln)
		*/
		fmt.Printf("%d \t %b \t %#x - %x \n", 13, 13, 13, 13)
	}

}
