//Channel test for synchronize the goroutines
package main

import (
	"fmt"
	"time"
)

var c chan int

func main() {
	c = make(chan int)
	go job("study", 2)
	go job("work", 3)
	fmt.Println("I'm waiting, but don't keep me waiting too long.")
	<-c
	<-c
}

func job(job string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(job, "is ready!")
	c <- 1
}
