package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// var c []byte
	fmt.Println("File Handling Here We Go..!!")
	file, e := os.Open("exmp.txt")
	if e == nil {
		fmt.Println("File Created")
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			wordList := strings.Fields(line)
			fmt.Println(wordList)
		}
	} else {
		fmt.Println(e)
	}
}
