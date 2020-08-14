// helloworld project main.go
package main

import (
	"fmt"
)

func main() {
    s1 := "helloworld"
    s2 := "hellotom"

    switch Compare(s1, s2) {
    case 0:
        fmt.Println("s1 equal to s2")
    case 1:
        fmt.Println("s1 greater than s2")
    case -1:
        fmt.Println("s1 less than s2")
    }
}

func Compare(s1, s2 string) int {
    var i int
    for i = 0; i < len(s1) && i < len(s2); i++ {
        switch {
        case s1[i] > s2[i]:
            return 1
        case s1[i] < s2[i]:
            return -1
        } 
    }

    switch {
    case i < len(s1):
        return 1
    case i < len(s2):
        return -1
    }

    return 0
}

