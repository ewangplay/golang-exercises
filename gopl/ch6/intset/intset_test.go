package intset

import (
	_ "fmt"
	"testing"
)

var s, s2 IntSet

//Test Add method
func TestAdd(t *testing.T) {
	if s.Len() != 0 {
		t.Fail()
	}

	s.Add(1)

	if s.Has(1) == false {
		t.Fail()
	}

	if s.Len() != 1 {
		t.Fail()
	}
}

//Test Remove method
func TestRemove(t *testing.T) {
	s.Remove(1)

	if s.Has(1) == true {
		t.Fail()
	}

	if s.Len() != 0 {
		t.Fail()
	}
}

//Test String method
func TestString(t *testing.T) {
	//Because the String() method's receiver is pointer to IntSet type,
	//so here must get address of the IntSet object s
	s.Add(1)
	s.Add(9)
	s.Add(144)

	//fmt.Printf("Set: %v\n", &s)
	t.Logf("Set s: %v\n", &s)
}

//Test Unionwith method
func TestUnionWith(t *testing.T) {
	s2.Add(9)
	s2.Add(50)
	s2.Add(6)
	s2.Add(200)

	t.Logf("Set s2: %v\n", &s2)

	s.UnionWith(&s2)

	if s.Len() != 6 {
		t.Fail()
	}

	t.Logf("Set s: %v\n", &s)
}

//Test Copy method
func TestCopy(t *testing.T) {
	s3 := s.Copy()

	if s3.Len() != 6 {
		t.Fail()
	}

	t.Logf("Set s3: %v\n", s3)
}

//Test Clear method
func TestClear(t *testing.T) {
	s.Clear()

	t.Logf("Set s:%v\n", &s)

	if s.Len() != 0 {
		t.Fail()
	}
}

//Test AddAll method
func TestAddAll(t *testing.T) {
	s.AddAll(299, 55, 2, 54, 10, 55, 69)

	t.Logf("Set s:%v\n", &s)

	if s.Len() != 6 {
		t.Fail()
	}
}

//Test IntersectWith method
func TestIntersectWith(t *testing.T) {
    s2.Clear()
	s2.AddAll(9, 50, 6, 200, 10, 55)

	t.Logf("Set s2: %v\n", &s2)

	s.IntersectWith(&s2)

	if s.Len() != 2 {
		t.Fail()
	}

	t.Logf("Set s: %v\n", &s)
}

//Test DifferenceWith method
func TestDifferenceWith(t *testing.T) {
    s.Clear()
	s.AddAll(299, 55, 2, 54, 10, 55, 69)
	s.DifferenceWith(&s2)

	if s.Len() != 4 {
		t.Fail()
	}

	t.Logf("Set s: %v\n", &s)
}
