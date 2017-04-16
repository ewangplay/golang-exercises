package intset

import (
    "testing"
    "fmt"
)

func TestAdd(t *testing.T) {
    var s IntSet

    fmt.Printf("Set Num: %d\n", s.Len())

    s.Add(1)
    if s.Has(1) == false {
        t.Fail()
    }

    fmt.Printf("Set Num: %d\n", s.Len())

    s.Remove(1)
    if s.Has(1) == true {
        t.Fail()
    }

    fmt.Printf("Set Num: %d\n", s.Len())
}

func TestHas(t *testing.T) {
    var s IntSet
    if s.Has(1) == true {
        t.Fail()
    }
}

 
