package intset

import (
	"bytes"
	"fmt"
)

const UINT_SIZE = 32<<(^uint(0)>>63)

//An IntSet is a set of small non-negative integers
//Its zero value represents the empty set
type IntSet struct {
	words []uint
}

//Has reports whether the set contains the integer value x
func (s *IntSet) Has(x int) bool {
	index, bit := x/UINT_SIZE, x%UINT_SIZE
	if len(s.words) > 0 {
		return index < len(s.words) && s.words[index]&(1<<uint(bit)) != 0
	}
	return false
}

//Add adds the specified integer x to the set
func (s *IntSet) Add(x int) {
	index, bit := x/UINT_SIZE, x%UINT_SIZE
	for index >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[index] |= 1 << uint(bit)
}

//Remove removes the integer x from the set
func (s *IntSet) Remove(x int) {
	index, bit := x/UINT_SIZE, x%UINT_SIZE
	if index < len(s.words) {
		s.words[index] &= ^(1 << uint(bit))
	}
}

//Len returns the number of elements
func (s *IntSet) Len() int {
	var sum int

	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				sum += 1
			}
		}
	}

	return sum
}

//Unionwith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	var i int
	for ; i < len(s.words) && i < len(t.words); i++ {
		s.words[i] |= t.words[i]
	}
	if i < len(t.words) {
		s.words = append(s.words, t.words[i:]...)
	}
}

//String returns the set as a string of form "{1, 2, 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer

	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", i*UINT_SIZE+j)
			}
		}
	}

	buf.WriteByte('}')

	return buf.String()
}

//Clear removes all the elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

//Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = append(t.words, s.words...)
	return &t
}

//Addall adds a list of values, such as s.Addall(1, 2, 3)
func (s *IntSet) AddAll(values ...int) {
	for _, v := range values {
		s.Add(v)
	}
}

//Intersectwith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	var new_set []uint

	for i := 0; i < len(s.words) && i < len(t.words); i++ {
		if s.words[i] == 0 || t.words[i] == 0 {
			continue
		}

		for j := 0; j < UINT_SIZE; j++ {
			if s.words[i]&(1<<uint(j)) != 0 && t.words[i]&(1<<uint(j)) != 0 {
				for i >= len(new_set) {
					new_set = append(new_set, 0)
				}
				new_set[i] |= 1 << uint(j)
			}
		}
	}

	s.words = new_set
}

//Differencewith sets s to the difference of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {

	for i := 0; i < len(s.words) && i < len(t.words); i++ {
		if s.words[i] == 0 || t.words[i] == 0 {
			continue
		}

		for j := 0; j < UINT_SIZE; j++ {
			if s.words[i]&(1<<uint(j)) != 0 && t.words[i]&(1<<uint(j)) != 0 {
				s.words[i] &= ^(1 << uint(j))
			}
		}
	}
}

//Symmetricdifferencewith sets to the symmetric difference of s and t
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	s1 := s.Copy()
	s1.DifferenceWith(t)

	s2 := t.Copy()
	s2.DifferenceWith(s)

	s1.UnionWith(s2)

	s.words = s1.words
}

//Elems returns the values list of s
func (s *IntSet) Elems() []int {
    var elems []int

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINT_SIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, i*UINT_SIZE+j)
			}
		}
	}

    return elems
}
