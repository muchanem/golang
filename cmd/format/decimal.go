package main

/*
  Printing numbers in decimal
*/
import (
	"fmt"
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {
	var number = input.Ask(c.Orange + "Please input a number: ")
	fmt.Println(c.Red + number)
	// just printing a decimal number in go using fmt
}
