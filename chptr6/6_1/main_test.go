package main

import (
	"math/rand"
	"os"
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

func TestHas(t *testing.T) {
	s := &IntSet{}
	r1 := rand.Uint64() % 64
	s.Add(int(r1))

	for i := 0; i < 64; i++ {
		if s.Has(i) && i != int(r1) {
			t.Errorf("has check not working %d %d", i, int(r1))
		}
	}

}

func TestUnion(t *testing.T) {
	s1 := &IntSet{}
	rand.Seed(int64(os.Getpid()))

	r1 := rand.Uint64() % 64
	s1.Add(int(r1))

	s2 := &IntSet{}
	s2.Add(165)

	s1.UnionWith(s2)

	if !s1.Has(int(r1)) {
		t.Errorf("expected value missing")
	}

	if !s1.Has(165) {
		t.Errorf("expected value")
	}
}

func TestString(t *testing.T) {

	s1 := &IntSet{}
	s2 := &IntSet{}

	s1.Add(1)
	s1.Add(2)
	s2.Add(164)
	s1.UnionWith(s2)
	str := s1.String()
	ss := "{1 2 164}"
	if str != ss {
		t.Errorf("unexpected return string: %s  expected %s", str, ss)
	}
}

func TestLen(t *testing.T) {

	s1 := &IntSet{}
	if s1.Len() != 0 {
		t.Errorf("before any adds, length should be 0")
	}

	s2 := &IntSet{}

	s1.Add(1)
	s1.Add(2)
	s2.Add(164)
	s1.UnionWith(s2)
	res := s1.Len()
	if res != 3 {
		t.Errorf("at this point Len=3, %d reported\n", res)
	}

	s1.Add(2)
	if res != 3 {
		t.Errorf("at this point Len=3, %d reported\n", res)
	}

}

func TestClear(t *testing.T) {
	s1 := &IntSet{}

	s1.Add(4)
	s1.Add(56)
	s1.Clear()
	if s1.Has(4) || s1.Has(56) {
		t.Errorf("set should be cleared")
	}
}
