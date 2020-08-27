/*
Methods can be declared only on named types (like Point) and pointers
to them (*Point), but thanks to embedding, it’s possible and sometimes
useful for unnamed struct types to have methods too.

This sample illustrates how to use unnamed struct type embedded methods.
*/
package main

import (
	"fmt"
	"sync"
)

// Defines a package-level variable using an unnamed struct type
// the sync.Mutex field is embedded within the unnamed struct type,
// its Lock and Unlock methods are promoted to the unnamed struct type,
// allowing us to lock the cache with a self-explanatory syntax.
var cache = struct {
	sync.RWMutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

// Lookup lookups the key in the cache and returns the value
func Lookup(key string) string {
	cache.RLock()
	v := cache.mapping[key]
	cache.RUnlock()
	return v
}

// Set sets the key / value pair into the cache
func Set(key, value string) {
	cache.Lock()
	cache.mapping[key] = value
	cache.Unlock()
}

func main() {
	Set("key1", "value1")
	fmt.Println(Lookup("key1"))
}
