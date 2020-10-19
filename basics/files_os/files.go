package main

import (
	"fmt"
	"os"
)

func readPartOfAfile() {
	barr := make([]byte, 100)
	f, err1 := os.Open("./basics/files_os/a.txt")
	nb, err2 := f.Read(barr)
	f.Close()
	fmt.Println(string(barr), nb, err1, err2)
}
func readEntireFile() {

}
func writeToAFile() {
	f, _ := os.Create("./basics/files_os/a.txt")
	barr := []byte("Blah Blah Blah")
	f.Write(barr)
	f.WriteString("Written using WriteString method")
}
func main() {
	readPartOfAfile()
}
