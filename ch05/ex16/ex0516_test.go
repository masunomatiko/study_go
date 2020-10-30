package ex0516

import "testing"

func Test_join(t *testing.T) {
	actual := join(",", "Alice", "Bob", "Carol")
	expected := "Alice,Bob,Carol"
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
