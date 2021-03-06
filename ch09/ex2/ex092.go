package popcount

import "sync"

var pc [256]byte
var loadTableOnce sync.Once

func PopCount(x uint64) int {
	pc := Table()
	return int(pc[byte(x>>(0*8))] + pc[byte(x>>(1*8))] + pc[byte(x>>(2*8))] + pc[byte(x>>(3*8))] + pc[byte(x>>(4*8))] + pc[byte(x>>(5*8))] + pc[byte(x>>(6*8))] + pc[byte(x>>(7*8))])
}

func Table() [256]byte {
	loadTableOnce.Do(loadTable)
	return pc
}

func loadTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
