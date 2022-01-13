package main

import "fmt"

// Functions are central in Go.

//sum

func plus(a int, b int) int {

	return a + b
}

func plusplus(a, b, c int) int {

	return a + b + c
}

func main() {
	ans := plus(1, 7)
	fmt.Println(ans)

	ans = plusplus(1, 4, 6)
	fmt.Println(ans)
}
