package main

import (
	"fmt"
	"time"
)

func main() {
	//seting time (0s for seconds, nanoseconds) utc for location
	time := time.Date(2017, time.January, 22, 540, 0, 0, time.UTC)
	fmt.Printf("When I wrote this program %s\n")
}
