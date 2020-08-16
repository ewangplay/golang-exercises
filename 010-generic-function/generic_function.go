/*
该示例演示了如何编写一个泛型函数。
*/
package main

import (
	"fmt"
)

// T 定义一个泛型类型
type T interface{}

func min(first T, rest ...T) T {
	var ret T
	switch first.(type) {
	case int:
		min := first.(int)
		for _, v := range rest {
			if i, ok := v.(int); ok {
				if i < min {
					min = i
				}
			} else {
				panic("type inconsistency")
			}
		}
		ret = min
	case string:
		min := first.(string)
		for _, v := range rest {
			if s, ok := v.(string); ok {
				if s < min {
					min = s
				}
			} else {
				panic("type inconsistency")
			}
        }
        ret = min
	}
	return ret
}

func main() {
	i := min(20, 33, 11, 44, 12, 25, 30, 15, 20)
	fmt.Printf("%T - %v\n", i, i)

	s := min("D", "B", "H", "A", "C", "G")
	fmt.Printf("%T - %v\n", s, s)
}
