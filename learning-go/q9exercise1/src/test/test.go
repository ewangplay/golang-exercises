package main

import (
    "container/stack"
    "fmt"
)

func main() {
    s := stack.NewIntStack(10)

    s.Push(2)
    s.Push(3)
    s.Push(5)

    fmt.Printf("My stack %v\n", s)

    if n1, ok := s.Pop(); ok {
        fmt.Println("pop data: ", n1)
    }

    if n2, ok := s.Pop(); ok {
        fmt.Println("pop data: ", n2)
    }

    fmt.Printf("My stack %v\n", s)
}
