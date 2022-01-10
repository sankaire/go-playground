package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	//set a vulue using arry index
	a[4] = 100
	fmt.Println(a)
	fmt.Println("len:", len(a))

	//declare an array and initalize in a single line

	d := [4]int{1, 2, 3, 4}
	fmt.Println(d)

	//two dimensional array
	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
