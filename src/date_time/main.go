package main

import (
	"fmt"
	"time"
)

func main() {
	//seting time (0s for seconds, nanoseconds) utc for GMT time, local for local time
	writeTime := time.Date(2017, time.January, 22, 9, 51, 0, 0, time.Local)
	fmt.Printf("When I wrote this program:%s\n", writeTime)
	// getting current time
	currrentTime := time.Now()
	// printing it as a string
	fmt.Printf("Current time:%s\n", currrentTime)
	// printing Month that I wrote the program
	fmt.Println("the Month is", currrentTime.Month())
	// printing the date I wrote the program
	fmt.Println("the date is", writeTime.Day())
	// printing the day I wrote the program
	fmt.Println("the day is", writeTime.Weekday())
	// defining tommorow
	//tommorrow := t.addDate()
}
