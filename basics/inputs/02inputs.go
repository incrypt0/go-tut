package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// fmt.Scanln(&name)
	// println(name)
	name := "Krish"
	fmt.Println(name)
	reader := bufio.NewReader(os.Stdin)
	myrating, _ := reader.ReadString('\n')
	fmt.Println(myrating)
}
