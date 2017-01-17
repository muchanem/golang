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

	println(strings.ToUpper(string1))
	println(strings.Title(string2))

}
