package sexpr

import "testing"

func TestFloat32(t *testing.T) {
	tests := []struct {
		input  float32
		expect string
	}{
		{1.5, "1.5"},
	}
	for _, test := range tests {
		data, err := Marshal(test.input)
		if err != nil {
			t.Errorf("Marshal(%f): %s", test.input, err)
		}
		if string(data) != test.expect {
			t.Errorf("Marshal(%f) actual %s, expect %s", test.input, data, test.expect)
		}
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		input  float64
		expect string
	}{
		{1.5, "1.5"},
	}
	for _, test := range tests {
		data, err := Marshal(test.input)
		if err != nil {
			t.Errorf("Marshal(%f): %s", test.input, err)
		}
		if string(data) != test.expect {
			t.Errorf("Marshal(%f) actual %s, expect %s", test.input, data, test.expect)
		}
	}
}

func TestComplex64(t *testing.T) {
	tests := []struct {
		input  complex64
		expect string
	}{
		{1 + 2i, "#C(1 2)"},
	}
	for _, test := range tests {
		data, err := Marshal(test.input)
		if err != nil {
			t.Errorf("Marshal(%v): %s", test.input, err)
		}
		if string(data) != test.expect {
			t.Errorf("Marshal(%v) actual %s, expect %s", test.input, data, test.expect)
		}
	}
}

func TestComplex128(t *testing.T) {
	tests := []struct {
		input  complex128
		expect string
	}{
		{1 + 2i, "#C(1 2)"},
	}
	for _, test := range tests {
		data, err := Marshal(test.input)
		if err != nil {
			t.Errorf("Marshal(%v): %s", test.input, err)
		}
		if string(data) != test.expect {
			t.Errorf("Marshal(%v) actual %s, expect %s", test.input, data, test.expect)
		}
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		input  bool
		expect string
	}{
		{true, "t"},
		{false, "nil"},
	}
	for _, test := range tests {
		data, err := Marshal(test.input)
		if err != nil {
			t.Errorf("Marshal(%v): %s", test.input, err)
		}
		if string(data) != test.expect {
			t.Errorf("Marshal(%v) got %s, wanted %s", test.input, data, test.expect)
		}
	}
}
