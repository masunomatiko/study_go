package tempconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCToK(t *testing.T) {
	x := AbsoluteZeroC
	actual := CtoK(x)
	expect := Kelvin(0)
	assert.Equal(t, actual, expect)
}

func TestKToC(t *testing.T) {
	x := Kelvin(0)
	actual := KtoC(x)
	expect := AbsoluteZeroC
	assert.Equal(t, actual, expect)
}
