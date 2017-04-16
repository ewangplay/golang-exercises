package main

import (
    "fmt"
)

func main() {
    for _,v := range fibonacci(80) {
        fmt.Println(v)
    }
}

func fibonacci(num int) (list []int64) {
    list = make([]int64, num)
    list[0] = 1; list[1] = 1
    for i:= 2; i<num; i++ {
        list[i] = list[i-1] + list[i-2]
    }
    return
}
