package main

import (
	"fmt"
	"time"

	c "github.com/skilstak/go-colors"
)

func main() {

	for {
		// sleep one second
		time.Sleep(100 * time.Millisecond)
		// clear screen
		fmt.Println(c.CL + "")
		// print time
		fmt.Printf(c.Multi("%s\n"), time.Now()))

	}
}
