/*
A bubble sort implementation with golang.
*/
package main

import (
    "fmt"
)

func main() {
    a := []int{20, 12, 34, 38, 16, 9, -1, -33, 43, 26, -15, 33}
    b := bubble_sort(a)

    fmt.Println("a: ", a)
    fmt.Println("b: ", b)
}

func bubble_sort(list []int) []int {
    BEGIN:
    swap_flag := false
    for i:=1; i<len(list)-1;i++ {
        if list[i-1] > list[i] {
            //利用golang语言中的平行赋值语法，可以大大简化列表前后两个
            //元素值交换的过程，这样的操作在C/C++语言中是不可行的，会导致
            //值覆盖。我推测golang语言在底层实现上使用了临时变量来缓存值，
            //但具体如何实现的，还需要查看golang的源代码才可以搞清楚。
            list[i-1],list[i] = list[i],list[i-1]
            swap_flag = true
        }
    }
    if swap_flag {
        goto BEGIN
    }
    return list
}

