package main

import "fmt"

func mayPanic() {
	panic("Critical eror")
}
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("after mayPanic()")
}
