package popcount

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
)

func ShiftPopCount(x uint8) int {
	var cnt int

	for i := 0; i < 8; i++ {
		cnt += int((x >> uint8(i)) & uint8(1))
	}

	return cnt
}

func DiffBitCount(x, y []byte) (int, error) {
	if len(x) != len(y) {
		return 0, nil
	}

	var cnt int
	for i := 0; i < len(x); i++ {
		cnt += ShiftPopCount(x[i] ^ y[i])

	}

	return cnt, nil
}

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	sha1 := sha256.Sum256([]byte(s1))
	sha2 := sha256.Sum256([]byte(s2))
	diff, err := DiffBitCount(sha1[:], sha2[:])
	if err != nil {
		fmt.Printf("%d", diff)
	} else {
		log.Fatal(err)
	}
}
