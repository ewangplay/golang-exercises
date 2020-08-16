/*
该例子用来演示如何动态创建函数的实现。
1. 我们首先声明一个package级别的变量，类型为函数的原型
2. 然后在package的init()函数中我们可以根据不同的条件来创建该函数不同的实现
3. 最后在main()函数我们来验证一下创建的动态函数是否按照预期来运行
*/
package main

import (
	"fmt"
	"os"
)

var myFunc func() string

func init() {
	if len(os.Args) > 1 {
		myFunc = func() string {
			return fmt.Sprintf("hello, %s", os.Args[1])
		}
	} else {
		myFunc = func() string {
			return "hello, world"
		}
	}
}

func main() {
	fmt.Println(myFunc())
}
