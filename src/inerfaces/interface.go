package main

import (
	"fmt"
	"math"
)

//basic interface for geometric shapes

type geometric interface {
	area() float64
	perim() float64
}
type rect struct {
	height, width float64
}
type circle struct {
	radius float64
}

//implemanttion of a rect
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//implemantation of a circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometric) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{height: 10, width: 5}
	c := circle{radius: 10}

	measure(r)
	measure(c)

}
