package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"
)

var logg *log.Logger
var wg sync.WaitGroup

func someHandler() {
	// Create the ctx with deadline
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))

	// Do some stuff
	wg.Add(1)
	go doStuff(ctx)

	// Wait for doStuff done
	wg.Wait()
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("ctx canceled: %v", ctx.Err())
			wg.Done()
			return
		default:
			logg.Printf("work...")
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("exit...")
}
