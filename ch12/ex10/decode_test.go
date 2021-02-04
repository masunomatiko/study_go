package decode

import (
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	tests := []struct {
		input  []byte
		expect bool
	}{
		{[]byte("t"), true},
		{[]byte("nil"), false},
	}
	var actual bool
	for _, test := range tests {
		if err := Unmarshal(test.input, &actual); err != nil {
			t.Errorf("Marshal(%s): %s", string(test.input), err)
		}
		if !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("not equal: actual=%v expect=%v", actual, test.expect)
		}
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		input  []byte
		expect float32
	}{
		{[]byte("1.5"), 1.5},
	}
	var actual float32
	for _, test := range tests {
		if err := Unmarshal(test.input, &actual); err != nil {
			t.Errorf("Marshal(%s): %s", string(test.input), err)
		}
		if !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("not equal: actual=%v expect=%v", actual, test.expect)
		}
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		input  []byte
		expect float64
	}{
		{[]byte("0.99"), 0.99},
	}
	var actual float64
	for _, test := range tests {
		if err := Unmarshal(test.input, &actual); err != nil {
			t.Errorf("Marshal(%s): %s", string(test.input), err)
		}
		if !reflect.DeepEqual(actual, test.expect) {
			t.Errorf("not equal: actual=%v expect=%v", actual, test.expect)
		}
	}
}
