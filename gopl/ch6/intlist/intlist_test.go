package intlist

import (
	"fmt"
	"testing"
)

func TestEmptyIntList(t *testing.T) {
	var list *IntList
	sum := list.Sum()
	if sum != 0 {
		t.Fatalf("sum got %v, want %v", sum, 0)
	}
}

func TestIntListAppend(t *testing.T) {
	list := &IntList{Value: 1}
	list = list.Append(&IntList{Value: 2})
	list = list.Append(&IntList{Value: 3})
	list = list.Append(&IntList{Value: 4})
	list = list.Append(&IntList{Value: 5})
	sum := list.Sum()
	if sum != 15 {
		t.Fatalf("sum got %v, want %v", sum, 15)
	}
}

func ExampleIntList_String_empty() {
	var list *IntList
	fmt.Println(list)
	// Output: []
}
func ExampleIntList_String_normal() {
	list := &IntList{Value: 1}
	list = list.Append(&IntList{Value: 2})
	list = list.Append(&IntList{Value: 3})
	list = list.Append(&IntList{Value: 4})
	list = list.Append(&IntList{Value: 5})
	fmt.Println(list)
	// Output: [1->2->3->4->5]
}
