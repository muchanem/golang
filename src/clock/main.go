package main

import (
	"fmt"
	"time"

	c "github.com/skilstak/go-colors"
)

func main() {

	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(c.CL + "")

		fmt.Printf(c.R+"%s\n", time.Now())

	}
}
