package main

import (
	"fmt"
	"reflect"
)

//In go we cannot simply pass an argument to a function without a type
//To do so we can use an empty interface
func typeChecker(thing interface{}) {
	fmt.Println(reflect.TypeOf(thing).String())
}
func main() {
	var a int
	typeChecker(a)
}
