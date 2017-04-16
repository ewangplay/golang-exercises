// Study Go -> Q6
package main

import (
	"fmt"
)

func main() {
    s := make([]float64, 0)
    //s := []float64{1.4, 2.34, 5.4, 7.443, 8.993, 6.123, 8845.243}
    r := average(s)
    fmt.Println("average: ", r)
}

func average(s []float64) (result float64) {
    //var sum float64
    sum := 0.0
    if len(s) == 0 {
        result = 0.0
    } else {
        for _,v := range s {
            sum += v
        }
        result = sum / float64(len(s))
    }
    return
}

