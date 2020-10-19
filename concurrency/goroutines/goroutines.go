package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	countTest5()
}
func countTest1() {
	//What happerns when we run this function only is that
	//There will be no output because
	//The goroutine runs in the background and the program execute the next statement
	//There is no next statement so it exits
	go count("Nairobi")
}
func countTest2() {
	//Here something interesting happens after mappping the goroutine to a thread
	//next statement is executed and therefore
	//We get a mixed output of both the goroutine and the
	//independant cloud function
	go count("Nairobi")
	count("Berlin")
}
func countTest3() {
	//Similar to 1 this one exists
	go count("Nairobi")
	go count("Berlin")
}
func countTest4() {
	//An easier way to continue this is by
	//Using a Scanln
	go count("Nairobi")
	go count("Berlin")
	fmt.Scanln()
}
func countTest5() {
	//This creates a wait group to wait for goroutines to complete

	var wg sync.WaitGroup
	wg.Add(1) //Adds a goroutine to counter

	go func() {
		count("Nairobi")
		wg.Done() //When the goroutine is done just decrement the counter
		//so that next piece of code can be executed
	}()
	wg.Wait()

}
func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
