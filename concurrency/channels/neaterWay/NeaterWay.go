package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go counter("berlin", c)

	for msg := range c {

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
