package main

import (
	"fmt"
)

func main() {
	// Creates a map with 10 capacity space
	m := make(map[string]string, 10)
	fmt.Println(len(m), m)

	// Adds two items to map
	m["a"] = "tom"
	m["b"] = "cap"
	fmt.Println(len(m), m)

	// Retrieves key 'a' with first synax
	v := m["a"]
	fmt.Println("a =", v)

	// Retrieves key 'b' with second synax
	v, found := m["b"]
	if found {
		fmt.Println("b =", v)
	} else {
		fmt.Println("key b not found")
	}

	// Retrieves a key that does not exist
	// Assigns the zero value of string type to v
	v = m["c"]
	fmt.Println("c =", v)

	// Retrieves a key that does not exist
	// Assigns the zero value of string type to v, false to found
	v, found = m["c"]
	if found {
		fmt.Println("c =", v)
	} else {
		fmt.Println("key c not found")
	}

	// Deletes one item from map
	delete(m, "b")

	fmt.Println(len(m), m)
}
