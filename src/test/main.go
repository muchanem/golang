package main

import (
	"time"

	c "github.com/skilstak/go-colors"
)

func main() {
	currentTime := time.now()
	printf(c.R+"The current time is: %s\n", currentTime)
}
