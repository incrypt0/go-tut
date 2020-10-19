package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go counter("Berlin  ", c)

	for {
		msg, open := <-c //here some sort of unpacking happens

		// ❗  The neater way is to use range c in the for statement itself
		//the second value returned by c is a boolean whether the channel is open
		//If the channel is not open then we break out of the channel
		if !open {
			break
		}
		fmt.Println(msg)
	}
}

//A channel is medium of transmission of data between
func counter(s string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- s //sending the string to the channel
		time.Sleep(time.Millisecond * 500)
	}

	close(c) //wwe can close the channel when we are done
}
