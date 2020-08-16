/*
该例子用来演示如何针对interfac{}类型的变量做类型确认(type assertion)
(1) resultOfType, boolean := expression.(Type) // Checked
(2) resultOfType := expression.(Type) // Unchecked; panic() on failure
*/
package main

import (
	"fmt"
)

func main() {
	var i interface{} = 99
	fmt.Printf("%T - %v\n", i, i)
	switch i.(type) {
	case int:
		v := i.(int)
		fmt.Printf("%T - %v\n", v, v)
	case int64:
		v := i.(int64)
		fmt.Printf("%T - %v\n", v, v)
	}

	i = 99.99
	if v, ok := i.(float32); ok {
		fmt.Printf("%T - %v\n", v, v)
	} else if v, ok := i.(float64); ok {
		fmt.Printf("%T - %v\n", v, v)
	} else {
		fmt.Printf("%T - %v\n", i, i)
	}

	i = []string{"tom", "jerry", "hello"}
	switch i.(type) {
	case string:
		v := i.(string)
		fmt.Printf("%T - %v\n", v, v)
	case []string:
		v := i.([]string)
		fmt.Printf("%T - %v\n", v, v)
	}
}
