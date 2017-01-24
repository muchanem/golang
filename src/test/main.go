package main

import (
	"fmt"
	"time"

	c "github.com/skilstak/go-colors"
)

func main() {

	var currentTime = time.now()

	fmt.Printf(c.R+"The current time is: %s\n", currentTime)
}
