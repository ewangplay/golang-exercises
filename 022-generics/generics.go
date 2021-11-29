package main

import (
	"fmt"
)

type stack[T comparable] struct{
	vars []T
}

func (s *stack[T])Push(t T) {
	s.vars = append(s.vars, t)
}

func (s *stack[T])Pop()(T, bool) {
	if len(s.vars) == 0 {
		var zero T
		return zero, false
	}
	top := s.vars[len(s.vars)-1]
	s.vars = s.vars[:len(s.vars)-1]
	return top, true
}

func (s *stack[T]) Contains(t T) bool {
	for _, v := range s.vars {
		if t == v {
			return true
		}
	}
	return false
}

func main() {
	var s stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)

	i, ok := s.Pop()
	if ok {
		fmt.Println("top value:", i)
	}

	if exists := s.Contains(i); exists {
		fmt.Println("there is", i)
	} else {
		fmt.Println("there is no", i)
	}
}
