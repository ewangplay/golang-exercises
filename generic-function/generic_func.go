package main

import (
    "fmt"
)

func Min(first interface{}, rest  ...interface{}) interface{} {
    min := first
    for _, v := range rest {
        switch v := v.(type) {
        case int:
            if v < min.(int) {
                min = v
            }
        case string:
            if v < min.(string) {
                min = v
            }
        }
    }
    return min
}

func Min1(first interface{}, rest  ...interface{}) interface{} {
    min := first
    switch first.(type) {
    case int:
        for _, v := range rest {
            if v.(int) < min.(int) {
                min = v
            }
        }
    case string:
        for _, v := range rest {
            if v.(string) < min.(string) {
                min = v
            }
        }
    }
    return min
}

func main() {
    i := Min(20, 33, 11, 8, 12, 25, 30, 15, 20)
    fmt.Printf("%T - %v\n", i, i)

    str := Min("D", "B", "H", "A", "C", "G")
    fmt.Printf("%T - %v\n", str, str)
}
