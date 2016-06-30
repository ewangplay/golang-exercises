/*
该例子用来演示如何动态创建函数的实现。
1. 我们首先声明一个package级别的变量，类型为函数的原型
2. 然后在package的init()函数中我们可以根据不同的条件来创建该函数不同的实现
3. 最后在main()函数我们来验证一下创建的动态函数是否按照预期来运行
(1) ./main tom
    > hello, Tom
(2) ./main jerry
    > hello, Jerry
(3) ./main other_str
    > hello, Nami
(4) ./main
    > hello, world
*/
package main

import (
    "os"
    "fmt"
)

var my_func func() string

func init() {
    if len(os.Args) > 1 {
        if os.Args[1] == "tom" {
            my_func = func() string {
                return "hello, Tom"
            }
        } else if os.Args[1] == "jerry" {
            my_func = func() string {
                return "hello, Jerry"
            }
        } else {
            my_func = func() string {
                return "hello, Nami"
            }
        }
    } else {
        my_func = func() string {
            return "hello, world"
        }
    }
}

func main() {
    fmt.Println(my_func())
}

