package main

import (
	"fmt"
	"sync"
)

type worker struct {
	Func func()
}

const (
	PoolSize = 5
)

func main() {
	var wg sync.WaitGroup

	channels := make(chan worker, 10)

	for i := 0; i < PoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for ch := range channels {
				ch.Func()
			}
		}()
	}

	for i := 0; i < 20; i++ {
		j := i //why to re-define one variable?
		wk := worker{
			Func: func() {
				fmt.Println(j)
			},
		}
		channels <- wk
	}

	close(channels)

	wg.Wait()
}
