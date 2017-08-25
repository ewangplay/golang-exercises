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
    var str interface{} = []string{"tom", "jerry", "hello"}

    if  i, ok := i.(int); ok {
        fmt.Printf("%T - %v\n", i, i)
    }

    if  i, ok := i.(int64); ok {
        fmt.Printf("%T - %v\n", i, i)
    }

    if  i, ok := i.(float32); ok {
        fmt.Printf("%T - %v\n", i, i)
    }

    if  i, ok := i.(float64); ok {
        fmt.Printf("%T - %v\n", i, i)
    }

    if  i, ok := i.(bool); ok {
        fmt.Printf("%T - %v\n", i, i)
    }

    if str, ok := str.(string); ok {
        fmt.Printf("%T - %v\n", str, str)
    }

    if str, ok := str.([]string); ok {
        fmt.Printf("%T - %v\n", str, str)
    }
}

