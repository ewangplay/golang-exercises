package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

func someHandler() {
	// Create the ctx
	ctx, cancel := context.WithCancel(context.Background())

	// Do some stuff
	go doStuff(ctx)

	// Cancel after 10 seconds
	time.Sleep(10 * time.Second)
	cancel()

}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("ctx canceled: %v", ctx.Err())
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
