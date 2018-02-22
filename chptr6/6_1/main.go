package intset

import (
	"bytes"
	"errors"
	"fmt"
)

//Type IntSet represents an in-memory set
type IntSet struct {
	words []uint64
}

//return a copy of the intset
func (s *IntSet) Copy() *IntSet {
	var n = IntSet{}
	copy(n.words, s.words)
	return &n
}

func (s *IntSet) Add(n int) error {
	if n < 0 {
		return errors.New("parameter must be non-negative")
	}

	//*increase the length of words if needed
	b := n / 64
	for b >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[b] |= (1 << uint(n%64))
	return nil
}

func (s *IntSet) Has(n int) bool {
	w := n / 64
	if w >= len(s.words) {
		return false
	}

	b := uint64(n % 64)
	return (s.words[w] & (1 << b)) != 0
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, wFromN := range t.words {
		if i < len(s.words) {
			s.words[i] |= t.words[i]
		} else {
			s.words = append(s.words, wFromN)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, j := range s.words {
		if j == 0 {
			continue
		}

		for k := 0; k < 64; k++ {
			if j&(1<<uint(k)) != 0 {
				if buf.Len() > 1 {
					buf.WriteString(" ")
				}
				fmt.Fprintf(&buf, "%d", 64*i+int(k))
			}
		}
	}
	buf.WriteString("}")

	return buf.String()
}

//Len returns the number of set elements
func (s *IntSet) Len() int {
	ret := 0
	for _, j := range s.words {
		if j == 0 {
			continue
		}

		for k := 0; k < 64; k++ {
			if j&(1<<uint(k)) != 0 {
				ret++
			}
		}
	}
	return ret
}

// Remove will clear set bit, this will not "shrink" the size of words even if all elements are cleared
func (s *IntSet) Remove(x int) error {
	w, b := x/64, uint64(x%64)

	if w > len(s.words) {
		return fmt.Errorf("Size of s.words could not contain %d", x)
	}

	targetBit := uint64(1 << b)
	if (s.words[w] & targetBit) == 0 {
		return fmt.Errorf("set did not contain %d", x)
	}
	s.words[w] &= ^targetBit
	return nil
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func main() {

	s := &IntSet{}
	s.Has(4)
}
