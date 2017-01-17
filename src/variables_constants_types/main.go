package main

import "strings"

func main() {
	string1 := "implicitly typed string"
	printf("str1: %v:%T\n", str1, str1)
	var string2 = "explicitly typed string"
	printf("str2: %v:%T\n", str2, str2)

	println(strings.ToUpper(string1))
	println(strings.Title(string2))

}
