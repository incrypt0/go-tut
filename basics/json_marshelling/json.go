package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string
	Email      string
	RollNo     int
	RegisterNo int
}

func main() {
	var p2 Person
	p1 := Person{
		Name:       "Joe",
		Email:      "joe@gmail.com",
		RollNo:     32,
		RegisterNo: 1232,
	}

	p1Json, _ := json.Marshal(p1)
	// s := string(p1Json)
	_ = json.Unmarshal(p1Json, &p2)
	fmt.Println(p2)
}
