package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := " So baby pull me closer\n In the back seat of your rover \n And i know you cant afford"
	scanner := bufio.NewScanner(strings.NewReader(s))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
