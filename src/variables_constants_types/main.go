package main

//
import (
	"fmt"
	"strings"
)

func main() {
	string1 := "implicitly typed string"
	fmt.Printf("string1: %v:%T\n", string1, string1)
	var string2 = "explicitly typed string"
	fmt.Printf("string2: %v:%T\n", string2, string2)
	// Puts string upper case
	fmt.Println(strings.ToUpper(string1))
	// Captializes every first letter
	fmt.Println(strings.Title(string2))
	lowerValue := "hello"
	upperValue := "HELLO"
	// == to see if they afre equal
	fmt.Println("Same?", (lowerValue == upperValue))
	fmt.Println("Equal when non case sensetive?", strings.EqualFold(lowerValue, upperValue))
	fmt.Println("Contains exp?", strings.Contains(string1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(string2, "exp"))

}
