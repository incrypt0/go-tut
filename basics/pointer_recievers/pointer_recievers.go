package main

import "fmt"

const usixteenbitmax float64 = 65535
const kmhmultiple float64 = 1.609

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	maxSpeed      float64
}

//If we want to modify a value inside of struct using a method
//we have to use pointers
func (myCar *car) speedUpgrade(newSpeed float64) {
	myCar.maxSpeed = newSpeed
}

func main() {
	a := car{gasPedal: 12, brakePedal: 1, steeringWheel: 1241, maxSpeed: 150}
	a.speedUpgrade(789)
	fmt.Println(a.maxSpeed)
}
