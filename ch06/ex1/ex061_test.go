package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Add(2000)
	actual := s.Len()
	expected := 2
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestRemove(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Add(2000)

	s.Remove(0)
	actual := s.Has(0)

	expected := false
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestClear(t *testing.T) {
	s := &IntSet{}
	s.Add(0)
	s.Add(1000)
	s.Clear()
	actual := s.Has(0) || s.Has(1000)
	expected := false
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestCopy(t *testing.T) {
	s := &IntSet{}
	s.Add(1)
	copy := s.Copy()
	copy.Add(2)
	actual := copy.Has(1) || copy.Has(2) || s.Has(2)
	expected := true
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
