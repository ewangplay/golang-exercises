package stack_test

import (
    "stack"
    "testing"
)

func TestIntStack(t *testing.T) {
    s := stack.NewIntStack(10)

    if ! s.Push(4) {
        t.Log("push 4 fail!")
        t.Fail()
    }

    num, ok := s.Pop()
    if ! ok {
        t.Log("pop data fail!")
        t.FailNow()
    }

    if num != 4 {
        t.Log("pop data incorrect!")
        t.FailNow()
    }
}

