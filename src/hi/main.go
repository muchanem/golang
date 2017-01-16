package main

import (
	c "github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {
	name := input.Ask(c.Red + "What is your name? ")
	println(c.Orange + "Hello World " + c.Green + name)
}
