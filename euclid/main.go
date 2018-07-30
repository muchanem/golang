package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/skilstak/go-input"
)

/*
  A ridiculously simple program that loops through the pythagrean triples
*/

func main() {
	input := input.Ask("Pythogream True of False, Pythgoream Loop, Pyhagoream Find (P, Pl, Pf)")
	input1 := strings.ToLower(input)
	switch input1 {
	case "p":
		preP()
	case "pl":
		pythagreanLoop()
	case "pf":
		pythagreanFind()
	}

}
func preP() {
	for m, n := 2, 1; m < math.MaxInt64; n, m = n+1, m+1 {
		time.Sleep(20000)
		fmt.Println(m, n)

	}
	time.Sleep(20000)
}
func pythagrean(m int, n int) {
	a := n ^ 2 - n ^ 2
	b := 2 * m * n
	c := m ^ 2 + n ^ 2
	fmt.Println(a, b, c)
	time.Sleep(100)
}
func pythagreanLoop() {

}
func pythagreanFind() {

}
