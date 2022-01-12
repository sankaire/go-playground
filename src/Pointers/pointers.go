package main

import (
	"fmt"
)

func zereVal(ival int) {
	ival = 0
}

func zeroPtr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)
	zereVal(i)
	fmt.Println("Zeroval:", i)
	zeroPtr(&i)
	fmt.Println("zeroptr:", i)
}
