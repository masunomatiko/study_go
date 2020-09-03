package popcount

func ClearPopCount(x uint64) int {
	var cnt int

	for x > 0 {
		x &= (x - uint64(1))
		cnt++
	}

	return cnt
}
