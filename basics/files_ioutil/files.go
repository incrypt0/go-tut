package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, e := ioutil.ReadFile("/home/incrypto/dev/go/src/basics/files_ioutil/abcd.txt")
	fmt.Println(string(dat), e)
	inp := []byte("Hello")
	err := ioutil.WriteFile("./basics/files_ioutil/blah.txt", inp, 0777)
	fmt.Println(err)

}
