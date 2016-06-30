package main

import (
	"fmt"
)

func main() {
    a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    fmt.Println(a)

    s1 := a[2:5]    //[3, 4, 5]
    fmt.Println(s1)

    s2 := a[:]      //[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
    fmt.Println(s2)

    s3 := a[:4]     //[1, 2, 3, 4]
    fmt.Println(s3)

    s4 := a[5:]     //[6, 7, 8, 9, 10]
    fmt.Println(s4)

    s5 := append(s4, 11, 12, 13)    //[6, 7, 8, 9, 10, 11, 12, 13
    fmt.Println(s5)

    s6 := append(s2, 11, 12, 13)    //[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]
    fmt.Println(s6)

    s7 := make([]int, 5)
    n1 := copy(s7, s6)
    fmt.Println("copy num: ", n1)   //copy num: 5
    fmt.Println(s7)                 //[1, 2, 3, 4, 5]

    n2 := copy(s7, s7[2:])
    fmt.Println("copy num: ", n2)   //copy num: 3
    fmt.Println(s7)                 //[3, 4, 5, 4, 5]
}
