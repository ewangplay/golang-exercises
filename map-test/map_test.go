package main

import (
	"fmt"
)

func main() {
    //create a map with 10 capacity space
	m1 := make(map[string]string, 10)

	fmt.Println(len(m1), m1)

    //add two items to map
	m1["a"] = "tom"
	m1["b"] = "cap"

	fmt.Println(len(m1), m1)

    v1 := m1["a"] 
    fmt.Println("a =", v1)

    v2, found := m1["b"]
    if found {
        fmt.Println("b =", v2)
    } else {
        fmt.Println("not found value for b")
    }

    v3 := m1["c"]
    fmt.Println("c =", v3)

    v4, found := m1["d"]
    if found {
        fmt.Println("d =", v4)
    } else {
        fmt.Println("not found value for d")
    }

    //delete one item from map
	delete(m1, "b")

	fmt.Println(len(m1), m1)
}
