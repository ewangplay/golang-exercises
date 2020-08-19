package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger
var ch = make(chan bool, 1)

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("exit...")
}

/*
context.WithCancel:

The returned context's Done channel is closed when the 
returned cancel function is called or when the parent 
context's Done channel is closed, whichever happens first.
*/
func someHandler() {
	// Create the ctx
	ctx, cancel := context.WithCancel(context.Background())

	// Do some stuff
	go doStuff(ctx)

	// Cancel after 10 seconds
	time.Sleep(10 * time.Second)
	cancel()

	// Waiting for doStuff to finish
	<-ch
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("ctx canceled: %v", ctx.Err())

			// Notify the main goroutin that I have finished
			ch <- true

			return
		default:
			logg.Printf("work...")
		}
	}
}
