package main

import (
	"fmt"
	"time"

	c "github.com/skilstak/go-colors"
)

func main() {

	var nowTime = time.Now()
	for true {
		fmt.Printf(c.R+"The current time is: %s\n", nowTime)

	}
}
