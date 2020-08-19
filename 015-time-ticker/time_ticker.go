/*
这段代码是time package中给出的官方例子，但是有问题: "time out"的场景始终不会触发。

select {
case m := <-c:
	handle(m)
case <-time.After(10 * time.Second):
	fmt.Println("timed out")
}

为什么time.After会不起作用呢？我们看一下time.After的源码就会明白：

func After(d Duration) <-chan Time {
	return NewTimer(d).C
}

要想触发Timer，After函数必须被调用，但在上面的例子中time.After一直没有机会被调用，也就
不能触发Timer。

修改成下面示例中的样子，Timer就可以正常触发，"time out"的场景就会出现。
*/
package main

import (
	"fmt"
	"time"
)

func ticker() {
	/*
		Tick is a convenience wrapper for NewTicker providing access to
		the ticking channel only. While Tick is useful for clients that
		have no need to shut down the Ticker, be aware that without a way
		to shut it down the underlying Ticker cannot be recovered by the
		garbage collector; it "leaks".
		注意：使用time.Tick是会造成资源泄露，因为底层的Ticker对象并不会被关闭和回收
		所以，请尽量使用NewTicker
	*/
	c1 := time.Tick(2 * time.Second)

	/*
		After waits for the duration to elapse and then sends the current
		time on the returned channel. It is equivalent to NewTimer(d).C.
		The underlying Timer is not recovered by the garbage collector
		until the timer fires. If efficiency is a concern, use NewTimer
		instead and call Timer.Stop if the timer is no longer needed.
		注意：使用time.After是会造成资源泄露，因为底层的Timer对象并不会被关闭和回收.
		所以，请尽量使用NewTimer
	*/
	c2 := time.After(10 * time.Second)

	for {
		select {
		case t := <-c1:
			fmt.Printf("ticker...%v\n", t)
		case t := <-c2:
			fmt.Printf("timed out...%v\n", t)
			return
		}
	}
}

func tickerImprove() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("ticker...%v\n", t)
		case t := <-timer.C:
			fmt.Printf("timed out...%v\n", t)
			return
		}
	}
}

func main() {
	ticker()
	tickerImprove()
}
