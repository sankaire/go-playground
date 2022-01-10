package main

import (
	"fmt"
	"math"
)

//const declares a constant value
const s string = "constant"

func main() {
	const k = 5000000
	const n = 3e20 / k
	fmt.Println(k)
	fmt.Println(int64(k))
	fmt.Println(math.Sin(k))
	// var declares more than 1 variable
	var a = 1
	fmt.Println(a)

	var b, c int = 3, 4
	fmt.Println(b, c)

	//variables with no initialization is 0
	var d int
	fmt.Println(d)

	// The := syntax is shorthand for declaring and initializing a variable
	e := "peter"
	fmt.Println(e)
}
