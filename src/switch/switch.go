package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//you can use comas to separate case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is not weekend")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
