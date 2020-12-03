package palindrome

import (
	"sort"
	"testing"
)

func TestPalindrome_奇数個(t *testing.T) {
	ints := []int{1, 2, 3, 4, 3, 2, 1}
	actual := IsPalindrome(sort.IntSlice(ints))
	expected := true
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestPalindrome_偶数個(t *testing.T) {
	ints := []int{1, 2, 3, 3, 2, 1}
	actual := IsPalindrome(sort.IntSlice(ints))
	expected := true
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func TestPalindrome_Fail(t *testing.T) {
	ints := []int{1, 2, 3, 3, 2, 2}
	actual := IsPalindrome(sort.IntSlice(ints))
	expected := false
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
