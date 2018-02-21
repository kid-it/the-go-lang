package main

import (
	"errors"
)

//Type IntSet represents an in-memory set
type IntSet struct {
	words []uint64
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
	if w > len(s.words) {
		return false
	}

	b := uint64(n % 64)
	return (s.words[w] & (1 << b)) != 0
}

func main() {

}
