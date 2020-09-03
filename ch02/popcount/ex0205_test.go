package popcount

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoopPopCount(t *testing.T) {
	var tests = []struct {
		x      uint64
		expect int
	}{
		{uint64(0), 0},
		{uint64(3), 2},
		{uint64(4), 1},
		{uint64(63), 6},
	}

	for _, test := range tests {
		actual := LoopPopCount(test.x)
		assert.Equal(t, actual, test.expect)
	}
}

func BenchmarkClearPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClearPopCount(uint64(i))
	}
}
