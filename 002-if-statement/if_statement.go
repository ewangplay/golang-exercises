package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(3, 5))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
