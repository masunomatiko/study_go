package intset

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	s := &IntSet{}
	s.AddAll(0, 2, 4)
	u := &IntSet{}
	u.AddAll(1, 2, 3)
	s.IntersectWith(u)
	if !s.Has(2) || s.Len() != 1 {
		t.Log(s)
		t.Fail()
	}
}

func TestDifferenceWith(t *testing.T) {
	s := &IntSet{}
	s.AddAll(0, 2, 4)
	u := &IntSet{}
	u.AddAll(1, 2, 3)
	s.DifferenceWith(u)
	expected := &IntSet{}
	expected.AddAll(0, 4)
	if s.String() != expected.String() {
		t.Log(s)
		t.Fail()
	}
}

func TestSymmetricDifference(t *testing.T) {
	s := &IntSet{}
	s.AddAll(0, 2, 4)
	u := &IntSet{}
	u.AddAll(1, 2, 3)
	s.SymmetricDifference(u)
	expected := &IntSet{}
	expected.AddAll(0, 1, 3, 4)
	if s.String() != expected.String() {
		t.Log(s)
		t.Fail()
	}
}
