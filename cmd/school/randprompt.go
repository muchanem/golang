package main

import "fmt"
import "time"
import "math/rand"

func main() {
	var prompts = []string{"Family", "Routine", "Schedule", "House"}
	rand.Seed(time.Now().UnixNano()) // takes the current time in nanoseconds as the seed
	n := rand.Intn(4)                // this gives you an int up to but not including 4
	fmt.Println(prompts[n])
}
