package intset

import "testing"

func TestAddAll(t *testing.T) {
	s := &IntSet{}
	s.AddAll(0, 100, 200, 300)
	actual := s.Has(0) && s.Has(100) && s.Has(200)
	expected := true
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
