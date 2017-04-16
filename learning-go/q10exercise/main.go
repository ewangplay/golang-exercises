/*
实现一个接受整数类型变参的函数，并且每行打印一个数字。
*/
package main

import (
    "fmt"
)

func main() {
    varargs_fun("arg list: ", 1, 3, 5, 7, 9)
    varargs_fun("这是一个关于变参的例子，变参列表: ", 2, 4, 6, 8,  19, 16, 19, 20)
}

func varargs_fun(prefix string, args ...int) {
    fmt.Println(prefix)
    //变参列表在函数中作为slice来处理
    for _,v := range args {
        fmt.Println(v)
    }
}
