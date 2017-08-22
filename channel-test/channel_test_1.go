//Channel test for transfer data between multi-goroutines
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	pending := make(chan int, 1)

	go func(ch chan int) {
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}(ch)

	var (
		i  int
		ok bool
	)

	reading := ch

	for {
		select {
		case i = <-pending:
			fmt.Printf("%d\n", i)
			reading = ch
		case i, ok = <-reading:
			if !ok {
				fmt.Printf("channel ch closed\n")
				return
			}
			pending <- i
			reading = nil
		}
	}

	fmt.Println("Exit...")

}
