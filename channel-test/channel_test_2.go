package main

import (
	"fmt"
)

/*
1. channel关闭后，仍然可以使用'<-'操作符读取该channel，读取到的值是定义该channel时数据类型的默认值
比如本例中定义chan使用的类型为interface{}，其默认值为nil

2. channel关闭后，不能再往该channel中写入值，会引起panic: send on closed channel
*/
func main() {
	ch := make(chan interface{})

	go func() {
		ch <- 1
		ch <- true
		ch <- "hello"
		ch <- 1.01
		close(ch)
	}()

	for {
		select {
		case v, ok := <-ch:
			if ok {
				fmt.Printf("v = %v\n", v)
			} else {
				fmt.Printf("Channel closed. Exit...\n")
				return
			}

			/*
				if v != nil {
					fmt.Printf("v = %v\n", v)
				} else {
					fmt.Printf("Channel closed. Exit...\n")
					// panic: send on closed channel
					// ch <- 1
					return
				}
			*/
		}
	}
}
