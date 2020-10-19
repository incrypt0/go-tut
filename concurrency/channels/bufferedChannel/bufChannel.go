package main

import "fmt"

func main() {
	c := make(chan string, 2) //Here we are specifying a capacity for the channel
	//When the channel recieves more than this capacity it pauses and looks for reciever
	c <- "💟"
	c <- "💙"
	//when we sent one more again deadlock happens
	msg := <-c

	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}
