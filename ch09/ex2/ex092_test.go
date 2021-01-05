package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(uint64(i))
	}
}

func BenchmarkPopCountOrigin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OriginalPopCount(uint64(i))
	}
}

// func init() {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// }

func LoopPopCount(x uint64) int {
	var cnt int

	for i := 0; i < 8; i++ {
		cnt += int(pc[byte(x>>(uint(i)*8))])
	}

	return cnt
}

func OriginalPopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] + pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] + pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
}
