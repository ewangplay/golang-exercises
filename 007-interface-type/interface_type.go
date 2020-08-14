package main

import (
	"fmt"
)

// I represents an interface
type I interface {
	Get() int
	Set(int)
}

// P implements the interface I
type P struct{ i int }

// Get gets value
func (p *P) Get() int {
	return p.i
}

// Set sets value
func (p *P) Set(num int) {
	p.i = num
}

func fn(t I) {
	fmt.Println(t.Get())
	t.Set(2)
	fmt.Println(t.Get())
}

func main() {
	p := P{1}
	fn(&p)
}
