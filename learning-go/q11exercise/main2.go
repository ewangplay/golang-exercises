package main

import (
    "fmt"
)

func main() {
    var i int64
    for i = 0; i < 80; i++ {
        fmt.Println(fibonacci(i))
    }
}

func fibonacci(num int64) int64 {
    if num == 0 || num == 1 {
        return 1
    }
    return fibonacci(num - 1) + fibonacci(num - 2)
}
