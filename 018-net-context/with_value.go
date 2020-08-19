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

/*
The provided key must be comparable and should not be of type string
or any other built-in type to avoid collisions between packages using context.

Users of WithValue should define their own types for keys.

To avoid allocating when assigning to an interface{}, context keys often have concrete type struct{}.
Alternatively, exported context key variables' static type should be a pointer or interface.
*/
type myContextKey string

const (
	UserTokenKey myContextKey = "User-Token"
	APIKey       myContextKey = "API-Key"
)

func main() {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("exit...")
}

/*
Use context Values only for request-scoped data that transits processes and APIs,
not for passing optional parameters to functions.
*/
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
	wg.Add(1)
	go doStuff(ctx)

	// Cancel after 10 seconds
	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logg.Printf("ctx canceled: %v", ctx.Err())
			wg.Done()
			return
		case <-ticker.C:
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
