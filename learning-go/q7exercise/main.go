// Study Go -> Q7
package main

import (
	"fmt"
)

func main() {
    fmt.Println(order(2, 7))
    fmt.Println(order(7, 2))
}

func order(n1, n2 int) (int, int) {
    if n1 < n2 {
        return n1, n2
    }
    return n2, n1
}
