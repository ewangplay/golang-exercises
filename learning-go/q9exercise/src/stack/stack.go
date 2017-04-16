/*
    Implementaion for int stack
    Interface:
        NewIntStack(max int)
        Push(elem int)
        Pop()
*/
package stack

import (
    "fmt"
)

type IntStack struct {
    data []int
    cap int
    pos int
}

func NewIntStack(max int) *IntStack {
    s := new(IntStack)
    s.data = make([]int, max)
    s.cap = max; s.pos = 0
    return s
}

func (s *IntStack) Push(elem int) (succ bool) {
    if s.pos < s.cap {
        s.data[s.pos] = elem
        s.pos++
        succ = true
        return
    }
    succ = false
    return
}

func (s *IntStack) Pop() (num int, succ bool) {
    if s.pos > 0 {
        s.pos--
        num = s.data[s.pos]
        s.data[s.pos] = 0
        succ = true
        return
    }
    succ = false
    return
}

func (s *IntStack) String() (fmt_str string) {
    for i := 0; i < s.pos ; i++ {
        fmt_str += fmt.Sprintf("[%d:%d] ", i, s.data[i])
    }
    return
}

