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

//methods
func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.maxSpeed / usixteenbitmax)
}

func main() {
	a := car{gasPedal: 12, brakePedal: 1, steeringWheel: 1241, maxSpeed: 150}
	b := new(car)
	fmt.Println(a.maxSpeed)
	fmt.Println(a.kmh(), b)
}
