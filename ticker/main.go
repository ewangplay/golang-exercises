package main

import (
	"fmt"
	"time"
)

/*
这段代码是time package中给出的官方例子，但是有问题
"time out"的场景始终不会触发。

select {
case m := <-c:
	handle(m)
case <-time.After(10 * time.Second):
	fmt.Println("timed out")
}

要想正确触发，需要把"time.After"预定定义成对象再使用，
就像下面代码中最终演示那样。
*/

func main() {
	c1 := time.Tick(2 * time.Second)
	c2 := time.After(10 * time.Second)

	for {
		select {
		case <-c1:
			fmt.Println("ticker...")
		case <-c2:
			fmt.Println("timed out, will exist...")
			return
		}
	}
}
