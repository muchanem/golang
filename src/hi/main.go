package main

import (
	"fmt"

	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {
	var name = input.Ask(c.Red + "What is your name? ")
	fmt.Println(c.Orange + "Hello World " + c.Green + name)
}
