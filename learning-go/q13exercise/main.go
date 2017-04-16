package main

import (
    "fmt"
)

const (
    FIND_MAX = 0
    FIND_MIN = 1
)

func main() {
    //a := []int{2, 4, 7, 10, 34, 25, 16, 9, 33, 45, 22}
    a := []int{-2, -4, -7, -10, -34, -25, -16, -9, -33, -45, -22}

    fmt.Println("list: ", a)
    fmt.Println("max value: ", find_max(a))
    fmt.Println("min value: ", find_min(a))

    if max,err := find_value(a, FIND_MAX); err == nil {
        fmt.Println("max value: ", max)
    }
    if min,err := find_value(a, FIND_MIN); err == nil {
        fmt.Println("min value: ", min)
    }
    if _,err := find_value(a, 3); err != nil {
        fmt.Println("error: ", err)
    }

}

func find_max(list []int) int {
    max_num := list[0]
    for _,value := range list {
        if max_num < value {
            max_num = value
        }
    }
    return max_num
}

func find_min(list []int) int {
    min_num := list[0]
    for _,value := range list {
        if min_num > value {
            min_num = value
        }
    }
    return min_num
}

func find_value(list []int, selector int) (int, error) {
    switch selector {
    case FIND_MAX:
        return find_max(list),nil
    case FIND_MIN:
        return find_min(list),nil
    }
    return 0, fmt.Errorf("invalid selector: %d", selector)
}

