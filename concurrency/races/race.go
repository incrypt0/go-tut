package main

// A race conditon occurs when two concurrent processes try to access same memory
// An example is a server handling an http request and increment a counter with no of requests
// If the server is handling requests concurrently at some point in time the memory location of the counter will get simultaneously written or accessed for reading
// At this point race condition occurs

//Here is an example for a race condition
import (
	"fmt"
	"time"
)

var count int

func race() {
	count++
}

//It can only occur if we increase the number of goroutines in this function
func main() {
	go race()
	go race()
	time.Sleep(1 * time.Second)
	fmt.Println(count)
}
