package main

import (
	"fmt"
	"os"
	"sort"
)

func isAnagram(s1, s2 string) bool {
	a1 := []rune(s1)
	a2 := []rune(s2)

	if len(a1) != len(a2) {
		return false
	}

	a1 = sortRuneArray(a1)
	a2 = sortRuneArray(a2)

	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}

func sortRuneArray(a []rune) []rune {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	return a
}

func main() {
	s1, s2 := os.Args[1], os.Args[2]
	fmt.Println(isAnagram(s1, s2))
}
