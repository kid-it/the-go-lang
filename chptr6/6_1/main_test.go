package main

import (
	"math/rand"
	"testing"
)

func Testadd(t *testing.T) {
	s := &IntSet{}
	r1 := rand.Uint64() % 64
	r2 := rand.Uint64() % 64
	s.Add(int(r1))
	s.Add(int(r2))

	if len(s.words) != 1 {
		t.Errorf("unexpected words length")
	}

	if s.words[0] != (1<<r1)|(1<<r2) {
		t.Errorf("incorrect result")
	}
}

func TestLargeadd(t *testing.T) {
	s := &IntSet{}
	r1 := rand.Uint64() % 64
	s.Add(64 + int(r1))

	if s.words[1] != (1 << r1) {
		t.Errorf("incorrect result %d", s.words[1])
	}

	r1 = rand.Uint64() % 64
	s.Add(1024 + int(r1))

	if s.words[16] != (1 << r1) {
		t.Errorf("incorrect result %d", s.words[16])
	}

}
