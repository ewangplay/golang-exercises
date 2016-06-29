package main

import (
    "fmt"
    "os"
)

func main() {
    a := max(3, 5)
    fmt.Println(a)

    b := readfile("1.txt")
    fmt.Println("read file ", b)
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func readfile(name string) bool {
    if _, err := os.Open(name); err != nil {
        return false
    }

    return true
}

