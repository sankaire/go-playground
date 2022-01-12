package main

import (
	"fmt"
)

//person struct
type person struct {
	name   string
	age    int
	isMale bool
}

//create a new person
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// p.isMale = true
	return &p
}

func main() {
	fmt.Println(person{"bob", 40, true})
	fmt.Println(person{"Peter", 21, true})
}
