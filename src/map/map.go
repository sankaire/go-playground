package main

/*--------------

Maps are Go’s built-in associative data type
(sometimes called hashes or dicts in other languages).

To create an empty map, use the builtin
make: make(map[key-type]val-type).

----------------*/

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	//Set key/value pairs using typical name[key] = val syntax.

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//get the value of the key with the name [key]

	v1 := m["k1"]

	fmt.Println("v1", v1)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	//The builtin delete removes key/value pairs from a map.
	delete(m, "k1")
	fmt.Println(m)

	//This can be used to disambiguate between missing keys and keys with zero values like 0 or ""
	_, psr := m["k1"]
	fmt.Println("psr", psr) //false

	//You can also declare and initialize a new map in the same line with this syntax.
	k := map[string]int{"foo": 1, "baa": 2}
	fmt.Println(k)
}
