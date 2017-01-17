package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "implicitly typed string"
	fmt.Printf("string1: %v:%T\n", string1, string1)
	var string2 = "explicitly typed string"
	fmt.Printf("string2: %v:%T\n", string2, string2)

	fmt.Println(strings.ToUpper(string1))
	fmt.Println(strings.Title(string2))
	lowerValue := "hello"
	upperValue := "HELLO"
	fmt.Println("Equal?", (lowerValue == upperValue))
	fmt.Println("Equal Non-Case-Sensitive?", strings.EqualFold(lowerValue, upperValue))

	fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(str2, "exp"))

}
