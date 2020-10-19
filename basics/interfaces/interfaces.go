package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	peri() float64
}
type rect struct {
	length, width float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
func (r rect) area() float64 {
	return r.length * r.width
}
func (c circle) peri() float64 {
	return math.Pi * 2 * c.radius
}
func (r rect) peri() float64 {
	return (r.length + r.width) * 2
}
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.peri())
}
func main() {
	rectangle := rect{
		length: 4,
		width:  5,
	}
	cir := circle{
		radius: 5,
	}
	measure(rectangle)
	measure(cir)
	fmt.Println("Hey there")
}
