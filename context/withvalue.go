package main

import (
	"context"
	"log"
	"os"
	"time"
)

var logg *log.Logger

const (
	UserTokenKey = "User-Token"
	APIKey       = "API-Key"
)

func someHandler() {
	// Create the ctx
	ctx, cancel := context.WithCancel(context.Background())

	// Add user token to ctx
	userToken := "12345678abc"
	ctx = context.WithValue(ctx, UserTokenKey, userToken)

	// Add API key to ctx
	apiKey := "somethingcrzyhaah901"
	ctx = context.WithValue(ctx, APIKey, apiKey)

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
			if userToken := ctx.Value(UserTokenKey); userToken != nil {
				logg.Printf("User Token: %v", userToken)
			}
			if apiKey := ctx.Value(APIKey); apiKey != nil {
				logg.Printf("API Key: %v", apiKey)
			}
		}
	}
}

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("exit...")
}
