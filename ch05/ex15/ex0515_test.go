package main

import "testing"

func Test_min(t *testing.T) {
	actual := min(3, -1, 4)
	expected := -1
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_minWithValidation(t *testing.T) {
	actual, err := minWithValidation(3, -1, 4)
	if err != nil {
		t.Errorf("something went wrong: %v", err)

	}
	expected := -1
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_max(t *testing.T) {
	actual := max(3, -1, 4)
	expected := 4
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_maxWithValidation(t *testing.T) {
	actual, err := maxWithValidation(3, -1, 4)
	if err != nil {
		t.Errorf("something went wrong: %v", err)

	}
	expected := 4
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}
