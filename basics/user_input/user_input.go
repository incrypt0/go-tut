package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Scan method inside fmt
func usingFmtScan(name string) {

	fmt.Println("Hey there, Enter your name ?")
	fmt.Scan(&name)
	fmt.Println(name)

}

//Scanln
func usingFmtScanln(name string) {

	fmt.Println("Hey there, Enter your name ?")
	fmt.Scanln(&name)
	fmt.Println(name)
}
func usingBufioReader(name string) {
	//Using bufio package
	//ReadLine method
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hey there, Enter your Full Name : ")
	blah, _, _ := reader.ReadLine()
	fmt.Println(string(blah))
}
func usingBufioReadString(name string) {
	//Using the ReadString method
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Hey there, Enter your Full Name : ")
	name, _ = reader.ReadString('\n') // this reads upto newline including newline
	name = strings.Replace(name, "\n", "", -1)
	fmt.Println(name)
	// fmt.Println('a')
}
func usingBufioScanner(name string) {
	//Scanner in bufio
	fmt.Print("Hey there, Enter your Full Name : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	fmt.Println(name)
}
func main() {
	var name string
	var choice int
	var loopIt bool = true
l1:
	for loopIt {
		fmt.Printf("1)Scan\n2)Scanln\n3)bufioReader\n4)bufioReadString\n5)bufioScanner\nEnter your choice :")
		fmt.Scanf("%d", &choice)
		switch choice {
		case 1:
			usingFmtScan(name)
		case 2:
			usingFmtScanln(name)

		case 3:
			usingBufioReader(name)
		case 4:
			usingBufioReadString(name)
		case 5:
			usingBufioScanner(name)
		default:
			break l1

		}
	}
}
