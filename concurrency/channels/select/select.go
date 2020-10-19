package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms 🔥"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2s 🔥"
			time.Sleep(time.Second * 2)
		}
	}()

	// ❗ Now the problem is that
	// When we use a simple for loop like this and just printout the values coming out of the channels
	// the first go routine will wait for a reciever and function goes to the next line and nex goroutines gets executed
	// then the second goroutine waits for the values
	// And at the reciever end when c1 is recieved and printed out c2 takes 2sec to recieve
	// Hence this makes goroutine 1 slower
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	//We can overcome this by using select statement
	//Select statement allows to recieve from whichever channel is ready to recive

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)

		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
