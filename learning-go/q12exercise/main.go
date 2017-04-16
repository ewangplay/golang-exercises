/*
Map function implementation.
    A map() -function is a function that takes a function and a list. The function is applied to each member in the list and a new list containing these calculated values is returned. Thus:
    map (f (), (a 1 , a 2 , . . . , a n−1 , a n )) = (f (a 1 ), f (a 2 ), . . . , f (a n−1 ), f (a n ))
*/
package main

import (
    "fmt"
)

func main() {
    a := []int{1, 2, 3, 4, 5}
    f := func (n int) int {
        return n * n
    }

    b := fnOperator(f, a)

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}

func fnOperator(fn func (int) int, list []int) []int {
    for i := 0; i < len(list); i++ {
        list[i] = fn(list[i])
    }
    return list
}
