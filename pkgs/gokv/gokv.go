package main

import (
	"fmt"

	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/badgerdb"
)

type foo struct {
	Bar string
}

func main() {
	options := badgerdb.DefaultOptions // Address: "localhost:6379", Password: "", DB: 0

	// Create store
	store, err := badgerdb.NewStore(options)
	if err != nil {
		panic(err)
	}
	defer store.Close()

	// Store, retrieve, print and delete a value
	interactWithStore(store)
}

// interactWithStore stores, retrieves, prints and deletes a value.
// It's completely independent of the store implementation.
func interactWithStore(store gokv.Store) {
	// Store value
	val := foo{
		Bar: "baz",
	}
	err := store.Set("foo123", val)
	if err != nil {
		panic(err)
	}

	// Retrieve value
	retrievedVal := new(foo)
	found, err := store.Get("foo123", retrievedVal)
	if err != nil {
		panic(err)
	}
	if !found {
		panic("Value not found")
	}

	fmt.Printf("foo: %+v\n", *retrievedVal) // Prints `foo: {Bar:baz}`

	// Delete value
	err = store.Delete("foo123")
	if err != nil {
		panic(err)
	}
}
