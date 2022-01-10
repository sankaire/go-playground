package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println(len(s))

	//append - new values
	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println(s)

	//slices can be copy'd
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
}
