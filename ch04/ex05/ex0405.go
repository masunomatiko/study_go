package main

import "fmt"

func remove_dups(a []string) []string {
	var idx int
	for _, s := range a {
		if a[idx] == s {
			continue
		}
		idx++
		a[idx] = s
	}
	return a[:idx+1]
}

func main() {
	a := []string{"a", "b", "c", "e", "d"}
	a = remove_dups(a)
	fmt.Println(a)
}
