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
	// printing it as a string current time as string using %s
	fmt.Printf("Current time:%s\n", currrentTime)
	// printing Month
	fmt.Println("the Month is", currrentTime.Month())
	// printing the date
	fmt.Println("the date is", currrentTime.Day())
	// printing the day
	fmt.Println("the day is", currrentTime.Weekday())
	// defining tommorow
	tommorrow := t.addDate(currrentTime.Year(), currrentTime.Month(), currrentTime.Day())
}
