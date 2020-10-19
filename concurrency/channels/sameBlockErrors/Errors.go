package main

import "fmt"

func main() {
	c := make(chan string)
	c <- "💥"
	msg := <-c
	fmt.Println(msg)

	// 😢 This one also leads to a deadlock
	// Reason : 👉 the entire logic is in the main functtion which is not  a goroutine
	// So when we sent data to a channel it waits until the data is recieved in another go routine usually
	// Since the entire logic is in this main funtion it starts waiting when we sent data across channel
	// Therefore the function will not move forward and deadlock occurs
}
