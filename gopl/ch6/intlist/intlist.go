/*
Package intlist implements a linked list of integers.

Just as some functions allow nil pointers as arguments, so do some
methods for their receiver, especially if nil is a meaningful zero
value of the type, as with maps and slices.

In this simple linked list of integers, nil represents the empty list
*/
package intlist

import (
	"fmt"
)

// An IntList is a linked list of integers
// An nil *IntList represents the empty list
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of list elements
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

// Append appends one integers list to current list
func (list *IntList) Append(l *IntList) *IntList {
	if list == nil {
		return l
	}
	tail := list
	for tail.Tail != nil {
		tail = tail.Tail
	}
	tail.Tail = l
	return list
}

// String returns human-readable string format of the linked list
func (list *IntList) String() string {
	if list == nil {
		return "[]"
	}
	s := "["
	p := list
	for p != nil {
		s += fmt.Sprintf("%v", p.Value)
		p = p.Tail
		if p != nil {
			s += "->"
		} else {
			s += "]"
		}
	}
	return s
}
