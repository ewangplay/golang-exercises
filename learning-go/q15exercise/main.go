package main

import (
    "fmt"
)

func main() {
    f1 := plusTwo()
    fmt.Printf("%v\n", f1(2))

    f2 := plusX(3)
    fmt.Println(f2(2))

    f3 := plusX(10)
    fmt.Println(f3(2))
}

func plusTwo () func (int) int {
    return func (num int) int {
        return num + 2
    }
}

func plusX(x int) func (int) int {
    //在这里我们创建了一个闭包(Closure)。
    //闭包一个匿名函数，并且在该匿名函数里可以“捕获”到它创建处可见的所有常量
    //变量，并且可以使用它们，而不用担心该闭包被调用时这些常量或变量超出生存
    //周期。
    //比如这个闭包，就使用本次变量x。
    return func (num int) int {
        return num + x
    }
}
