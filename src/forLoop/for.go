package main

import (
	"fmt"
)

func main() {
	i := 1
	//single condition loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//initial/condition/ after loop
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}
	//for without a condition will loop untill you break
	for {
		fmt.Println("loop")
		break
	}
	//you can continue to the next itteration
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
