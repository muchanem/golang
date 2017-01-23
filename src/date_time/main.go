package main

import (
	"fmt"
	"time"
)

func main() {
	//seting time (0s for seconds, nanoseconds) utc for GMT time, local for local time
	writeTime := time.Date(2017, time.January, 22, 9, 51, 0, 0, time.Local)
	fmt.Printf("When I wrote this program:%s\n", writeTime)

	currrentTime := time.Now()
	fmt.Printf("Current time:%.14f\n", currrentTime)

}
