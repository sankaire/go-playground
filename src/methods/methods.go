package main

import "fmt"

type rectangle struct {
	width, height int
}

//this area method has a receiver type of*rectangle
func (r *rectangle) area() int {
	return r.height * r.width
}

func (r *rectangle) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	r := rectangle{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r

	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
