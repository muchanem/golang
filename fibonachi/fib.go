package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/skilstak/go-colors"
	"github.com/skilstak/go-input"
)

func main() {

	n := input.Ask(colors.R + "Find Number or loop through (f or l)?")
	if n == "l" {
		fibnext()
	} else {
		a := input.Ask(colors.R + "Number?")
		g, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(err)
		}
		var t = float64(g)
		fibfind(t)
	}
}
func round(n float64) int {
	if n < 0 {
		return int(n - 0.5)
	}
	return int(n + 0.5)
}
func fibnext() int {
	for X := 0; X <= 90; X++ {
		var e = float64(X)
		fibfind(e)
	}
	X := 0
	return X
}
func fibfind(n float64) float64 {
	z := math.Phi
	h := 1 / math.Sqrt(5) * math.Pow(z, n)
	u := round(h)
	fmt.Println(u)

	return h
}
