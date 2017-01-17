package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implicitly typed string"
	Printf("str1: %v:%T\n", str1, str1)
	var str2 = "An explicitly typed string"
	Printf("str2: %v:%T\n", str2, str2)

	println(strings.ToUpper(str1))
	println(strings.Title(str2))

	lValue := "hello"
	uValue := "HELLO"
	fmt.Println("Equal?", (lValue == uValue))
	fmt.Println("Equal Non-Case-Sensitive?", strings.EqualFold(lValue, uValue))

	fmt.Println("Contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("Contains exp?", strings.Contains(str2, "exp"))
}
