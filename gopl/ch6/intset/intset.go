package intset

//An IntSet is a set of small non-negative integers
//Its zero value represents the empty set
type IntSet struct {
    words []uint64
}

//Has reports whether the set contains the integer value x
func (s *IntSet) Has(x int) bool {
    index, bit := x/64, x%64
    if len(s.words) > 0 {
        return index < len(s.words) && s.words[index] & (1<<uint(bit)) != 0
    }
    return false
}

//Add adds the specified integer x to the set
func (s *IntSet) Add(x int) {
    index, bit := x/64, x%64
    for index >= len(s.words) {
        s.words = append(s.words, 0)
    }
    s.words[index] |= 1 << uint(bit)
}

//Remove removes the integer x from the set
func (s *IntSet) Remove(x int) {
    index, bit := x/64, x%64
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
        for j := 0; j < 64; j++ {
            if word & (1<<uint(j)) != 0 {
                sum += 1
            }
        }
    }

    return sum
}

//Unionwith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
}

//String returns the set as a string of form "{1, 2, 3}" 
func (s *IntSet) String() string {
    return ""
}

//Clear removes all the elements from the set
func (s *IntSet) Clear() {
}

//Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
    return nil
}

