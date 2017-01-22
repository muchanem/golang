package main

import (
	"fmt"
	"time"
)

func main() {
	//seting time (0s for seconds, nanoseconds) utc for location
	time := time.Date(2017, time.January, 22, 9, 51, 0, 0, time.Local)
	fmt.Printf("When I wrote this program:%s\n", time)
}
