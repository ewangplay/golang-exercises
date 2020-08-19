/*
该示例用来演示channel关闭后的读写行为。

1. channel关闭后，仍然可以使用'<-'操作符读取该channel，
读取到的值是定义该channel数据类型的默认值，比如本例中
定义chan使用的类型为interface{}，其默认值为nil

2. channel关闭后，不能再往该channel中写入值，
会引起panic: "send on closed channel"
*/
package main

import (
	"fmt"
)

func main() {
	ch := make(chan interface{})

	go func() {
		ch <- 1
		ch <- true
		ch <- "hello"
		ch <- 1.01
		close(ch)
	}()

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%v\n", e)
		}
	}()

	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Printf("v = %v\n", v)
			} else {
				fmt.Printf("Channel closed. Exit...\n")

				// 关闭channel后仍然读取
				fmt.Println(<-ch) // nil

				// 关闭channel后仍然写入
				ch <- 1 // panic: send on closed channel
				return
			}
		}
	}
}
