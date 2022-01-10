package main

import (
	"fmt"
)

func main() {
	fmt.Println("values")
	fmt.Println("go" + "lang")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("6.2/2.0 =", 6.2/2.0)

	fmt.Println(true && false) //returns false
	fmt.Println(true || false) //returns true
	fmt.Println(!true)         //retuns false
}
