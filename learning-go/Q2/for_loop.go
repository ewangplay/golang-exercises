/* 
<<Learning Go>> exercises Q2:
For-loop:
1. Create a simple loop with the for construct. Make it loop 10 times and print out the loop counter with the fmt package.
2. Rewrite the loop from 1. to use goto. The keyword for may not be used.
3. Rewrite the loop again so that it fills an array and then prints that array to the screen.
*/
package main

import (
	"fmt"
)

func main() {
    //Q2-1
    func1()

    //Q2-2
    func2()

    //Q2-3
    func3()
}

func func1() {
    for i:=0; i<10; i++ {
        if i < 9 {
            fmt.Print(i, " ")
        } else {
            fmt.Println(i)
        }
    }
}

func func2() {
    i := 0
    BEGIN:
    if i < 9 {
        fmt.Print(i, " ")
    } else {
        fmt.Println(i)
    }

    i++
    if i<10 {
        goto BEGIN
    }
}

func func3() {
    var a [10]int

    for i := 0; i < len(a); i++ {
        a[i] = i
    }

    for _,v := range a {
        if v < 9 {
            fmt.Print(v, " ")
        } else {
            fmt.Println(v)
        }
    }
}

